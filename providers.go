package goefa

var Providers = map[string]EFAProvider{

	// Augsburger Verkehrsverbund
	"avv": EFAProvider{
		BaseURL:                  "http://efa.avv-augsburg.de/avv/",
		DepartureMonitorEndpoint: "XML_DM_REQUEST",
		StopFinderEndpoint:       "XML_STOP_FINDER_REQUEST",
		TripEndpoint:             "XML_TRIP_REQUEST",
	},

	// Münchner Verkehrs- und Tarifverbund
	"mvv": EFAProvider{
		BaseURL:                  "http://efa.mvv-muenchen.de/",
		DepartureMonitorEndpoint: "XML_DM_REQUEST",
		StopFinderEndpoint:       "XML_STOP_FINDER_REQUEST",
		TripEndpoint:             "XML_TRIP_REQUEST",
	},
}
