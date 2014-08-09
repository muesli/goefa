/*
 *    Copyright (C) 2014 Michael Wendland
 *
 *    This program is free software: you can redistribute it and/or modify
 *    it under the terms of the GNU Affero General Public License as published
 *    by the Free Software Foundation, either version 3 of the License, or
 *    (at your option) any later version.
 *
 *    This program is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU Affero General Public License for more details.
 *
 *    You should have received a copy of the GNU Affero General Public License
 *    along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 *    Authors:
 *      Michael Wendland <michiwend@michiwend.com>
 */

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
	Name string

	BaseURL                  string
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
