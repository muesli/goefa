/*
 * Copyright (C) 2018 Christian Muehlhaeuser
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
 *   Christian Muehlhaeuser <muesli@gmail.com>
 */

package goefa

import (
	"net/url"
	"strconv"
	"time"
)

// Trip represents a trip from one stop to another
type Trip struct {
	OriginID            int
	DestinationID       int
	OriginPlatform      string
	DestinationPlatform string
	DepartureTime       time.Time
	ArrivalTime         time.Time
	MeansOfTransport    MeansOfTransport
}

// Route is a route between two stops, including all necessary transfers
type Route struct {
	Trips         []*Trip
	OriginID      int
	DestinationID int
	DepartureTime time.Time
	ArrivalTime   time.Time
}

type tripResult struct {
	Stops []struct {
		ID           int     `xml:"stopID,attr"`
		Platform     string  `xml:"platform,attr"`
		PlatformName string  `xml:"platformName,attr"`
		Locality     string  `xml:"locality,attr"`
		Lat          float64 `xml:"x,attr"`
		Lng          float64 `xml:"y,attr"`
		DateTime     EFATime `xml:"itdDateTime"`
	} `xml:"itdPoint"`
	MeansOfTransport MeansOfTransport `xml:"itdMeansOfTransport"`
}

type routesResult struct {
	Trips []*tripResult `xml:"itdPartialRouteList>itdPartialRoute"`
}

type tripsResult struct {
	Response
	Odv struct {
		OdvPlace struct {
		}
		OdvName struct {
			State string `xml:"state,attr"`
		} `xml:"itdOdvName"`
	} `xml:"itdTripRequest>itdOdv"`
	Routes []*routesResult `xml:"itdTripRequest>itdItinerary>itdRouteList>itdRoute"`
}

func (d *tripsResult) endpoint() string {
	return "XML_TRIP_REQUEST2"
}

// Route requests an itinerary of trips from the stop with ID origin to the stop with ID destination
func (efa *Provider) Route(origin int, destination int, due time.Time) ([]*Route, error) {
	rt := "0"
	if efa.EnableRealtime {
		rt = "1"
	}

	params := url.Values{
		"locationServerActive": {"1"},
		"useRealtime":          {rt},
		"type_origin":          {"stopID"},
		"type_destination":     {"stopID"},
		"name_origin":          {strconv.FormatInt(int64(origin), 10)},
		"name_destination":     {strconv.FormatInt(int64(destination), 10)},
		"itdDate":              {due.Format("20060102")},
		"itdTime":              {due.Format("1504")},
	}

	var result tripsResult
	if err := efa.postRequest(&result, params); err != nil {
		return nil, err
	}

	var routes []*Route
	for _, route := range result.Routes {
		r := Route{
			OriginID:      origin,
			DestinationID: destination,
			DepartureTime: route.Trips[0].Stops[0].DateTime.Time,
			ArrivalTime:   route.Trips[len(route.Trips)-1].Stops[1].DateTime.Time,
		}

		for _, trip := range route.Trips {
			if len(trip.Stops) != 2 {
				panic("Did not expect more than two stops in a single trip")
			}
			t := Trip{
				DepartureTime:       trip.Stops[0].DateTime.Time,
				ArrivalTime:         trip.Stops[1].DateTime.Time,
				MeansOfTransport:    trip.MeansOfTransport,
				OriginPlatform:      replaceEmptyWithNA(trip.Stops[0].Platform),
				DestinationPlatform: replaceEmptyWithNA(trip.Stops[1].Platform),
				OriginID:            trip.Stops[0].ID,
				DestinationID:       trip.Stops[1].ID,
			}
			if t.MeansOfTransport.MotType == 0 && t.MeansOfTransport.Number == "" {
				t.MeansOfTransport.MotType = -1
			}

			r.Trips = append(r.Trips, &t)
		}

		routes = append(routes, &r)
	}

	return routes, nil
}
