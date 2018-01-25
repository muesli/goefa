/*
 * Copyright (C) 2014      Michael Wendland
 *               2014-2018 Christian Muehlhaeuser
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
 *   Christian Muehlhaeuser <muesli@gmail.com>
 */

package goefa

import (
	"net/url"
	"strconv"
	"time"
)

// Departure represents a single departure for a specific stop.
type Departure struct {
	CountDown    int         // minutes until departure
	MapLink      string      // link to map
	Platform     string      // platform code
	PlatformName string      // platform name
	StopID       int         // ID of the stop
	StopName     string      // name of the stop
	Lat          int64       // Latitude
	Lng          int64       // Longitude
	DateTime     time.Time   // Timestamp of departure
	ServingLine  ServingLine // Line
}

type departureResult struct {
	// Area         int    `xml:"area,attr"`
	CountDown    int    `xml:"countdown,attr"`
	MapName      string `xml:"mapName,attr"`
	Platform     string `xml:"platform,attr"`
	PlatformName string `xml:"platformName,attr"`
	StopID       int    `xml:"stopID,attr"`
	StopName     string `xml:"stopName,attr"`
	Lat          int64  `xml:"x,attr"`
	Lng          int64  `xml:"y,attr"`

	DateTime    EFATime     `xml:"itdDateTime"`
	ServingLine ServingLine `xml:"itdServingLine"`
}

type departureMonitorResult struct {
	Response
	Odv struct {
		OdvPlace struct {
		}
		OdvName struct {
			State string `xml:"state,attr"`
		} `xml:"itdOdvName"`
	} `xml:"itdDepartureMonitorRequest>itdOdv"`
	Departures []*departureResult `xml:"itdDepartureMonitorRequest>itdDepartureList>itdDeparture"`
}

func (d *departureMonitorResult) endpoint() string {
	return "XML_DM_REQUEST"
}

// Departures performs a stateless dm_request for the corresponding stopID and
// returns an array of Departures. Use time.Now() as the second argument in
// order to get the very next departures. The third argument determines how
// many results will be returned by EFA.
func (efa *Provider) Departures(stopID int, due time.Time, results int) ([]*Departure, error) {
	rt := "0"
	if efa.EnableRealtime {
		rt = "1"
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

	var deps []*Departure
	for _, dep := range result.Departures {
		d := Departure{
			CountDown:    dep.CountDown,
			MapLink:      dep.MapName,
			Platform:     dep.Platform,
			PlatformName: dep.PlatformName,
			StopID:       dep.StopID,
			StopName:     dep.StopName,
			Lat:          dep.Lat,
			Lng:          dep.Lng,
			DateTime:     dep.DateTime.Time,
			ServingLine:  dep.ServingLine,
		}
		deps = append(deps, &d)
	}

	return deps, nil
}
