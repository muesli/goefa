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
	"errors"
	"net/url"
	"strconv"
	"time"
)

// Stop represents a public transportation stop
type Stop struct {
	ID             int
	Name           string
	Locality       string
	Lat            float64
	Lng            float64
	IsTransferStop bool

	Provider *Provider
}

type stopResult struct {
	ID             int     `xml:"id,attr"`
	StopID         int     `xml:"stopID,attr"`
	Name           string  `xml:"objectName,attr"`
	ValueName      string  `xml:",chardata"`
	Locality       string  `xml:"locality,attr"`
	Lat            float64 `xml:"x,attr"`
	Lng            float64 `xml:"y,attr"`
	IsTransferStop bool    `xml:"isTransferStop,attr"`
}

type stopFinderResult struct {
	Response
	Odv struct {
		OdvPlace struct {
		}
		OdvName struct {
			State string        `xml:"state,attr"`
			Stops []*stopResult `xml:"odvNameElem"`
		} `xml:"itdOdvName"`
	} `xml:"itdStopFinderRequest>itdOdv"`
}

func (s *stopFinderResult) endpoint() string {
	return "XML_STOPFINDER_REQUEST"
}

// FindStop queries the EFA StopFinder API of the corresponding provider and
// returns an array of matched stops (or only the identified one) or an error
// in case somthing went wrong.
func (efa *Provider) FindStop(name string) ([]*Stop, error) {
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
		return nil, err
	}

	stops := []*Stop{}
	for _, stop := range result.Odv.OdvName.Stops {
		stops = append(stops, &Stop{
			ID:             stop.ID,
			Name:           stop.Name,
			Locality:       stop.Locality,
			Lat:            stop.Lat,
			Lng:            stop.Lng,
			IsTransferStop: stop.IsTransferStop,
			Provider:       efa,
		})
	}

	switch result.Odv.OdvName.State {
	case "identified":
		fallthrough
	case "list":
		return stops, nil
	default:
		return nil, errors.New("no matched stops")
	}
}

// Stop queries the EFA StopFinder API of the corresponding provider for a stop
// with a specific ID.
func (efa *Provider) Stop(id int) (*Stop, error) {
	params := url.Values{
		"type_sf":              {"stopID"},
		"name_sf":              {strconv.FormatInt(int64(id), 10)},
		"outputFormat":         {"XML"},
		"stateless":            {"1"},
		"locationServerActive": {"1"},
		"anyObjFilter_sf":      {"2"},
		"coordOutputFormat":    {"WGS84[DD.ddddd]"},
	}

	var result stopFinderResult
	if err := efa.postRequest(&result, params); err != nil {
		return nil, err
	}

	if len(result.Odv.OdvName.Stops) > 0 {
		switch result.Odv.OdvName.State {
		case "identified":
			stop := Stop{
				ID:             result.Odv.OdvName.Stops[0].ID,
				Name:           result.Odv.OdvName.Stops[0].ValueName,
				Locality:       result.Odv.OdvName.Stops[0].Locality,
				Lat:            result.Odv.OdvName.Stops[0].Lat,
				Lng:            result.Odv.OdvName.Stops[0].Lng,
				IsTransferStop: result.Odv.OdvName.Stops[0].IsTransferStop,
				Provider:       efa,
			}
			if stop.ID == 0 {
				stop.ID = result.Odv.OdvName.Stops[0].StopID
			}

			return &stop, nil
		}
	}
	return nil, errors.New("no matched stops")
}

// Departures is just a helper method and calls the Departures() method of the
// Provider.
func (stop *Stop) Departures(due time.Time, results int) ([]*Departure, error) {
	return stop.Provider.Departures(stop.ID, due, results)
}

// Route is just a helper method and calls the Route() method of the Provider.
func (stop *Stop) Route(destination int, due time.Time) ([]*Route, error) {
	return stop.Provider.Route(stop.ID, destination, due)
}
