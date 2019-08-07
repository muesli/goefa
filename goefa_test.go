package goefa

import (
	"testing"
	"time"
)

var (
	providers = []*Provider{
		NewProvider("http://efa.avv-augsburg.de/avv/", true),
		NewProvider("http://efa.mvv-muenchen.de/mvv/", true),
	}
)

func TestStops(t *testing.T) {
	tests := []struct {
		provider   *Provider
		stop       string
		expectedID int
	}{
		{providers[0], "Koenigsplatz", 2000101},
		{providers[0], "Kongresshalle", 2000114},
		{providers[0], "Zoo/Bot. Garden", 2000687},
		{providers[1], "Koenigsplatz", 1000110},
	}

	for _, tt := range tests {
		stop, err := tt.provider.Stop(tt.expectedID)
		if err != nil {
			t.Errorf("Stop not found %d: %s", tt.expectedID, err)
		}
		if stop.ID != tt.expectedID {
			t.Errorf("Expected Stop ID %d, got %d", tt.expectedID, stop.ID)
			return
		}

		stops, err := tt.provider.FindStop(tt.stop)
		if err != nil {
			t.Errorf("Stop not found %s: %s", tt.stop, err)
			return
		}
		if len(stops) != 1 {
			t.Errorf("Expected only one result, got %d", len(stops))
			return
		}
		if stops[0].ID != tt.expectedID {
			t.Errorf("Expected Stop ID %d, got %d", tt.expectedID, stops[0].ID)
			return
		}
	}
}

func TestDepartures(t *testing.T) {
	tests := []struct {
		provider *Provider
		stopID   int
		results  int
	}{
		{providers[0], 2000101, 3},
		{providers[1], 1000110, 3},
	}

	for _, tt := range tests {
		departures, err := tt.provider.Departures(tt.stopID, time.Now(), tt.results)
		if err != nil {
			t.Error(err)
			return
		}

		if len(departures) != tt.results {
			t.Errorf("Expected %d results, got %d", tt.results, len(departures))
		}

		/*
			for _, dep := range departures {
				fmt.Printf("%17s %-5s due in %-2d minutes (%s) --> %s\n",
					dep.ServingLine.MotType.String(),
					dep.ServingLine.Number,
					dep.Countdown,
					dep.DateTime.Format("15:04"),
					dep.ServingLine.Direction)
			}
		*/
	}
}

func TestTrips(t *testing.T) {
	tests := []struct {
		provider      *Provider
		originID      int
		destinationID int
	}{
		{providers[0], 2000114, 2000687},
	}

	for _, tt := range tests {
		routes, err := tt.provider.Route(tt.originID, tt.destinationID, time.Now())
		if err != nil {
			t.Error(err)
			return
		}

		if len(routes) < 1 {
			t.Error("Expected a route, but got none")
			return
		}

		/*
			for _, trip := range routes[0].Trips {
				fmt.Printf("%+v\n", trip)
			}
		*/
	}
}
