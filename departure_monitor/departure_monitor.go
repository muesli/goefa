package main

import (
	"flag"
	"fmt"

	"github.com/michiwend/goefa"
)

func main() {
	efa := *goefa.Providers["avv"]

	stop := flag.String("stop", "Königsplatz", "id or (part of the) stop name")
	max_results := flag.Int("results", 5, "how many results to show")
	flag.Parse()

	station, err := efa.FindStation(*stop)
	if err != nil {
		fmt.Println("Stop does not exist or name is not unique!")
		return
	}
	fmt.Printf("Selected stop: %s (%d)\n\n",
		station.IdfdStop.StopName,
		station.IdfdStop.StopID)

	departures, err := efa.Departures(station, *max_results)
	if err != nil {
		fmt.Println("Could not retrieve departure times!")
		return
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
