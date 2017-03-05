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
	"errors"
	"net/url"
	"time"
)

type EFAStop struct {
	Id             int     `xml:"id,attr"`
	Name           string  `xml:"objectName,attr"`
	Locality       string  `xml:"locality,attr"`
	Lat            float64 `xml:"x,attr"`
	Lng            float64 `xml:"y,attr"`
	IsTransferStop bool    `xml:"isTransferStop,attr"`

	Provider *EFAProvider
}

// Departures is just a helper method and calls the Departures() method of
// EFAProvider.
func (stop *EFAStop) Departures(due time.Time, results int) ([]*EFADeparture, error) {
	return stop.Provider.Departures(stop.Id, due, results)
}

type stopFinderResult struct {
	EFAResponse
	Odv struct {
		OdvPlace struct {
		}
		OdvName struct {
			State string     `xml:"state,attr"`
			Stops []*EFAStop `xml:"odvNameElem"`
		} `xml:"itdOdvName"`
	} `xml:"itdStopFinderRequest>itdOdv"`
}

func (s *stopFinderResult) endpoint() string {
	return "XML_STOPFINDER_REQUEST"
}

// FindStop queries the EFA StopFinder API of the corresponding provider and
// returns whether the stop was identified/unique (bool), an array of matched
// stops (or only the identified one) or an error in case somthing went wrong.
func (efa *EFAProvider) FindStop(name string) (bool, []*EFAStop, error) {

	// To get a more detailed response from the StopFinder request we can use
	// EFAs LocationServer (locationServerActive=1, type_sf=any). To limit the
	// results to a specific type we can use anyObjFilter_sf=<bitmask> as
	// following:
	//	0 any type
	//	1 locations
	//	2 stations
	//	4 streets
	//	8 addresses
	//	16 crossroads
	//	32 POIs
	//	64 postal codes
	// "stations and streets" results in 2 + 4 = 6
	params := url.Values{
		"type_sf":              {"any"},
		"name_sf":              {name},
		"outputFormat":         {"XML"},
		"stateless":            {"1"},
		"locationServerActive": {"1"},
		"anyObjFilter_sf":      {"2"},
		"coordOutputFormat":    {"WGS84[DD.ddddd]"},
	}

	var result stopFinderResult

	if err := efa.postRequest(&result, params); err != nil {
		return false, nil, err
	}

	for _, stop := range result.Odv.OdvName.Stops {
		stop.Provider = efa
	}

	switch result.Odv.OdvName.State {
	case "identified":
		return true, result.Odv.OdvName.Stops, nil
	case "list":
		return false, result.Odv.OdvName.Stops, nil
	default:
		return false, nil, errors.New("no matched stops")
	}

}
