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

	"net/http"
	"net/url"
	"strconv"
	"time"

	"code.google.com/p/go-charset/charset"
)

type EFAStop struct {
	Id             int    `xml:"id,attr"`
	Name           string `xml:"objectName,attr"`
	Locality       string `xml:"locality,attr"`
	IsTransferStop bool   `xml:"isTransferStop,attr"`
	Provider       *EFAProvider

	//FIXME include coordinates
}

// Departures performs a stateless dm_request for the corresponding EFAStop and
// returns an array of EFADepartures. Use time.Now() as the first argument in
// order to get the very next departures. The second argument determines how
// many results will be returned by EFA
func (stop *EFAStop) Departures(time time.Time, results int) (
	[]*EFADeparture, error) {

	params := url.Values{
		"type_dm":              {"any"},
		"name_dm":              {strconv.Itoa(stop.Id)},
		"locationServerActive": {"1"},
		"useRealtime":          {"1"},
		"dmLineSelection":      {"all"}, //FIXME enable line selection
		"limit":                {strconv.Itoa(results)},
		"mode":                 {"direct"},
		"stateless":            {"1"},
		"itdDate":              {time.Format("20060102")},
		"itdTime":              {time.Format("1504")},
	}

	resp, err := http.PostForm(
		stop.Provider.BaseURL+stop.Provider.DepartureMonitorEndpoint,
		params)
	defer resp.Body.Close()

	if err != nil {
		return []*EFADeparture{}, err
	}

	decoder := xml.NewDecoder(resp.Body)
	decoder.CharsetReader = charset.NewReader

	var result efaDepartureMonitorRequest

	if err = decoder.Decode(&result); err != nil {
		return []*EFADeparture{}, err
	}

	return result.Departures, nil

}
