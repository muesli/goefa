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

type EFA struct {
	BaseURL string
}

//FIXME: separate goefa structs (like Station) and XML structs
func (efa *EFA) FindStation(name string) (*StopInfo, error) {
	//FIXME: nicer impl: use station search api if avail
	endpoint := "XML_DM_REQUEST"
	params := url.Values{
		"type_dm":              {"stop"},
		"name_dm":              {name},
		"useRealtime":          {"1"},
		"locationServerActive": {"1"},
		"dmLineSelection":      {"all"},
		"limit":                {"5"},
		"mode":                 {"direct"},
	}

	resp, err := http.PostForm(efa.BaseURL+endpoint, params)
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
func (efa *EFA) Departures(station *StopInfo, results int) ([]Departure, error) {
	endpoint := "XML_DM_REQUEST"
	params := url.Values{
		"type_dm":              {"stop"},
		"name_dm":              {strconv.Itoa(station.IdfdStop.StopID)},
		"useRealtime":          {"1"},
		"locationServerActive": {"1"},
		"dmLineSelection":      {"all"},
		"limit":                {strconv.Itoa(results)},
		"mode":                 {"direct"},
	}

	resp, err := http.PostForm(efa.BaseURL+endpoint, params)
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
