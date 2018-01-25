/*
 * Copyright (C) 2014      Michael Wendland
 *               2014-2018 Christian Muehlhaeuser
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published
 * by the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 * Authors:
 *   Michael Wendland <michael@michiwend.com>
 *   Christian Muehlhaeuser <muesli@gmail.com>
 */

/*
Package goefa implements a go (golang) client library to access data of public
transport companies which provide an EFA interface. You can search a stop, get
its next departures or request a trip.
*/
package goefa

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/rogpeppe/go-charset/charset"
	_ "github.com/rogpeppe/go-charset/data"
)

type efaEndpoint interface {
	endpoint() string
}

func (efa *Provider) postRequest(result efaEndpoint, params url.Values) error {
	client := http.Client{}
	reqURL, err := url.Parse(efa.BaseURL)
	if err != nil {
		return err
	}
	reqURL.Path = path.Join(reqURL.Path, result.endpoint())

	req, err := http.NewRequest("POST", reqURL.String(), strings.NewReader(params.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set(
		"User-Agent",
		"GoEFA, a golang EFA client / 0.0.1 (https://github.com/muesli/goefa)",
	)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	defer req.Body.Close()

	b, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(b))

	decoder := xml.NewDecoder(bytes.NewReader(b))
	decoder.CharsetReader = charset.NewReader
	return decoder.Decode(result)
}

func replaceEmptyWithNA(s string) string {
	if strings.TrimSpace(s) == "" {
		return "n/a"
	}

	return s
}
