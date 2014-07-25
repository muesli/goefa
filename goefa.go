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

func (efa *EFA) FindStation(name string) (int, error) {
	//FIXME: impl eventually
	return -1, errors.New("Not yet implemented")
}

//FIXME: turn station_id into an int (or does EFA really use a string as id?)
func (efa *EFA) Departures(station_id string, results int) ([]Departure, error) {
	endpoint := "XML_DM_REQUEST"
	params := url.Values{
		"type_dm":              {"stop"},
		"name_dm":              {station_id},
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
