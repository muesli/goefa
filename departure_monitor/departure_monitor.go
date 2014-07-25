package main

import (
	"flag"
	"fmt"

	"github.com/muesli/goefa"
)

func main() {
	efa := goefa.EFA{}

	station_id := flag.String("stop", "Königsplatz", "id or (part of the) stop name")
	max_results := flag.Int("results", 5, "how many results to show")
	flag.StringVar(&efa.BaseURL, "baseurl", "http://efa.avv-augsburg.de/avv/", "base-url for EFA API")
	flag.Parse()

/*	if result.Stop.State != "identified" {
		fmt.Println("Stop does not exist or name is not unique!")
	}
	fmt.Printf("Selected stop: %s (%d)\n\n",
		result.Stop.IdfdStop.StopName,
		result.Stop.IdfdStop.StopID)*/

	departures, err := efa.Departures(*station_id, *max_results)
	if err != nil {
		fmt.Println("Stop does not exist or name is not unique!")
	}
	for _, departure := range departures {
		plu := " "
		if departure.Countdown != 1 {
			plu = "s"
		}

		fmt.Printf("Route %-5s due in %-2d minute%s --> %s\n",
			departure.ServingLine.Number,
			departure.Countdown,
			plu,
			departure.ServingLine.Direction)
	}
}
