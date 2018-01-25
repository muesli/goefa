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

package main

import (
	"errors"
	"flag"
	"fmt"
	"time"

	"github.com/muesli/goefa"
)

func pickStop(stops []*goefa.Stop) (*goefa.Stop, error) {
	if len(stops) > 1 {
		fmt.Println("Two or more stops where matched:")
		for i, stop := range stops {
			fmt.Printf("%2d - %s (%s)\n", i, stop.Name, stop.Locality)
		}

		fmt.Print("Choose one: ")
		var i int
		_, err := fmt.Scanf("%d", &i)
		if err != nil {
			return nil, err
		}

		if i >= len(stops) {
			return nil, errors.New("index out of range")
		}

		return stops[i], nil
	}

	fmt.Println("Stop identified: " + stops[0].Name)
	return stops[0], nil
}

func main() {
	pname := flag.String("provider", "http://efa.avv-augsburg.de/avv/", "Base URL of the EFA API")
	origin := flag.String("origin", "Haunstetten Nord", "Name of the origin stop")
	dest := flag.String("destination", "Zoo", "Name of the destination stop")
	flag.Parse()

	provider := goefa.NewProvider(*pname, true)
	origins, err := provider.FindStop(*origin)
	if err != nil {
		fmt.Println(err)
		return
	}
	dests, err := provider.FindStop(*dest)
	if err != nil {
		fmt.Println(err)
		return
	}

	oStop, err := pickStop(origins)
	if err != nil {
		fmt.Println(err)
		return
	}
	dStop, err := pickStop(dests)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("\nYour next departures:")
	routes, err := provider.Route(oStop.ID, dStop.ID, time.Now())
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, route := range routes {
		fmt.Printf("Trip duration: %s, %d transfers\n", route.ArrivalTime.Sub(route.DepartureTime), len(route.Trips)-1)
		for _, trip := range route.Trips {
			origin, err := provider.Stop(trip.OriginID)
			if err != nil {
				fmt.Println(err)
				return
			}
			dest, err := provider.Stop(trip.DestinationID)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Printf("%17s %-5s %s (%s, platform %s) --> %s (%s, platform %s)\n",
				trip.MeansOfTransport.MotType.String(),
				trip.MeansOfTransport.Number,
				origin.Name,
				trip.DepartureTime.Format("15:04"),
				trip.OriginPlatform,
				dest.Name,
				trip.ArrivalTime.Format("15:04"),
				trip.DestinationPlatform)
		}
		fmt.Println()
	}
}
