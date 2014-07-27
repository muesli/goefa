package goefa

import (
	"encoding/xml"
	"errors"
	_ "fmt"
	"net/http"
	"net/url"
	"strconv"

	"code.google.com/p/go-charset/charset"
	_ "code.google.com/p/go-charset/data"
)

type EFAProvider struct {
	BaseURL string

	DepartureMonitorEndpoint string
	StopFinderEndpoint       string
	TripEndpoint             string

	//FIXME: include general params for all requests (e.g. useRealtime, ...)
}

//FIXME: separate goefa structs (like Station) and XML structs
func (efa *EFAProvider) FindStop(name string) (*StopInfo, error) {
	// FindStop queries the stopfinder API and returns Stops matching 'name'

	params := url.Values{
		"type_sf":              {"stop"},
		"name_sf":              {name},
		"locationServerActive": {"1"},
		"outputFormat":         {"XML"},
		"stateless":            {"1"},
		// "limit":                {"5"},
		// "mode":                 {"direct"},
	}

	resp, err := http.PostForm(efa.BaseURL+efa.StopFinderEndpoint, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result StopResult
	decoder := xml.NewDecoder(resp.Body)
	decoder.CharsetReader = charset.NewReader
	if err = decoder.Decode(&result); err != nil {
		return nil, err
	}

	if result.Stop.State != "identified" {
		return nil, errors.New("Stop does not exist or name is not unique!")
	}

	return &result.Stop, nil
}

//FIXME: turn station_id into an int
func (efa *EFAProvider) Departures(station *StopInfo, results int) ([]Departure, error) {
	params := url.Values{
		"type_dm":              {"stop"},
		"name_dm":              {strconv.Itoa(station.IdfdStop.StopID)},
		"useRealtime":          {"1"},
		"locationServerActive": {"1"},
		"dmLineSelection":      {"all"},
		"limit":                {strconv.Itoa(results)},
		"mode":                 {"direct"},
	}

	resp, err := http.PostForm(efa.BaseURL+efa.DepartureMonitorEndpoint, params)
	if err != nil {
		return []Departure{}, err
	}
	defer resp.Body.Close()

	var result StopResult
	decoder := xml.NewDecoder(resp.Body)
	decoder.CharsetReader = charset.NewReader
	if err = decoder.Decode(&result); err != nil {
		return []Departure{}, err
	}
	//fmt.Printf("%+v", result)

	if result.Stop.State != "identified" {
		return []Departure{}, errors.New("Stop does not exist or name is not unique!")
	}

	return result.Departures, nil
}
