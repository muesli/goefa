/*
 * Copyright (C) 2014 Michael Wendland
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published
 * by the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 * Authors:
 *   Michael Wendland <michael@michiwend.com>
 */

package goefa

import (
	"net/url"
	"strconv"
	"time"
)

// EFADepartureArrival represents either an arrival or a departure and is not
// directly used but embedded by EFAArrival and EFADeparture.
type EFADepartureArrival struct {
	Area         int    `xml:"area,attr"`
	Countdown    int    `xml:"countdown,attr"`
	MapName      string `xml:"mapName,attr"`
	Platform     string `xml:"platform,attr"`
	PlatformName string `xml:"platformName,attr"`
	StopID       int    `xml:"displayName,attr"`
	StopName     string `xml:"stopName,attr"`
	Lat          int64  `xml:"x,attr"`
	Lng          int64  `xml:"y,attr"`

	DateTime    EFATime        `xml:"itdDateTime"`
	ServingLine EFAServingLine `xml:"itdServingLine"`
}

// EFAArrival represents a single arrival for a specific stop.
type EFAArrival struct {
	EFADepartureArrival `xml:"itdArrival"`
}

// EFADeparture represents a single departure for a specific stop.
type EFADeparture struct {
	EFADepartureArrival `xml:"itdDeparture"`
}

type departureMonitorResult struct {
	Odv struct {
		OdvPlace struct {
		}
		OdvName struct {
			State string `xml:"state,attr"`
		} `xml:"itdOdvName"`
	} `xml:"itdDepartureMonitorRequest>itdOdv"`

	Departures []*EFADeparture `xml:"itdDepartureMonitorRequest>itdDepartureList>itdDeparture"`
}

func (d *departureMonitorResult) endpoint() string {
	return "XML_DM_REQUEST"
}

func (efa *EFAProvider) Departures(stopID int, due time.Time, results int) ([]*EFADeparture, error) {
	var rt string

	if efa.EnableRealtime {
		rt = "1"
	} else {
		rt = "0"
	}

	params := url.Values{
		"type_dm":              {"any"},
		"name_dm":              {strconv.Itoa(stopID)},
		"locationServerActive": {"1"},
		"useRealtime":          {rt},
		"dmLineSelection":      {"all"}, //FIXME enable line selection
		"limit":                {strconv.Itoa(results)},
		"mode":                 {"direct"},
		"stateless":            {"1"},
		"itdDate":              {due.Format("20060102")},
		"itdTime":              {due.Format("1504")},
	}

	var result departureMonitorResult

	if err := efa.postRequest(&result, params); err != nil {
		return nil, err
	}

	return result.Departures, nil

}
