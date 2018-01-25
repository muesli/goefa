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

package main

import (
	"errors"
	"flag"
	"fmt"
	"time"

	"github.com/muesli/goefa"
)

func main() {
	pname := flag.String("provider", "http://efa.avv-augsburg.de/avv/", "Base URL of the EFA API")
	query := flag.String("stop", "Koenigsplatz", "The stop name to search for")
	results := flag.Int("results", 5, "How many results to show")
	flag.Parse()

	myprovider := goefa.NewProvider(*pname, true)
	stops, err := myprovider.FindStop(*query)
	if err != nil {
		fmt.Println(err)
		return
	}

	var mystop *goefa.Stop
	if len(stops) > 1 {
		fmt.Println("Two or more stops where matched:")
		for i, stop := range stops {
			fmt.Printf("%2d - %s (%s)\n", i, stop.Name, stop.Locality)
		}

		fmt.Print("Choose one: ")
		var i int
		_, err = fmt.Scanf("%d", &i)
		if err != nil {
			fmt.Println(err)
			return
		}

		if i >= len(stops) {
			fmt.Println(errors.New("index out of range"))
			return
		}

		mystop = stops[i]
	} else {
		mystop = stops[0]
		fmt.Println("Stop identified: " + mystop.Name)
	}
	fmt.Println("Your next departures:")

	departures, err := mystop.Departures(time.Now(), *results)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, dep := range departures {
		plu := " "
		if dep.CountDown != 1 {
			plu = "s"
		}

		destStop, err := myprovider.Stop(dep.ServingLine.DestStopID)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("%17s %-5s due in %-2d minute%s (%s, platform %s) --> %s\n",
			dep.ServingLine.MotType.String(),
			dep.ServingLine.Number,
			dep.CountDown,
			plu,
			dep.DateTime.Format("15:04"),
			dep.Platform,
			destStop.Name)
	}
}
