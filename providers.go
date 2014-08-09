package goefa

// Map of supported EFA providers
var Providers = map[string]EFAProvider{

	"avv": EFAProvider{
		Name:                     "Augsburger Verkehrs- und Tarifverbund GmbH",
		BaseURL:                  "http://efa.avv-augsburg.de/avv/",
		DepartureMonitorEndpoint: "XML_DM_REQUEST",
		StopFinderEndpoint:       "XML_STOPFINDER_REQUEST",
		TripEndpoint:             "XML_TRIP_REQUEST",
	},

	"mvv": EFAProvider{
		Name:                     "Münchner Verkehrs- und Tarifverbund",
		BaseURL:                  "http://efa.mvv-muenchen.de/mvv/",
		DepartureMonitorEndpoint: "XML_DM_REQUEST",
		StopFinderEndpoint:       "XML_STOPFINDER_REQUEST",
		TripEndpoint:             "XML_TRIP_REQUEST",
	},
}
