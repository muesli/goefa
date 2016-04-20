/*
 * Copyright (C) 2014 Michael Wendland
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
 */

/*
package goefa implements a go (golang) client library to access data of public
transport companies which provide an EFA interface. You can search a stop, get
its next departures or request a trip.
*/
package goefa

import (
	"encoding/xml"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/rogpeppe/go-charset/charset"
	_ "github.com/rogpeppe/go-charset/data"
)

// EFAProvider represents a public transport company that provides access to
// its EFA instance. Use providers.json to store a list of known providers.
type EFAProvider struct {
	Name           string
	BaseURL        string //FIXME use url.URL
	EnableRealtime bool
}

type efaResult interface {
	endpoint() string
}

type efaResponse struct {
	XMLName xml.Name `xml:"itdRequest"`

	client     string `xml:"client,attr"`
	clientIP   string `xml:"clientIP,attr"`
	language   string `xml:"language,attr"`
	lengthUnit string `xml:"lengthUnit,attr"`
	now        string `xml:"now,attr"`
	nowWD      int    `xml:"nowID,attr"`
	serverID   string `xml:"serverID,attr"`
	sessionID  int    `xml:"sessionID,attr"`
	version    string `xml:"version,attr"`
	virtDir    string `xml:"virtDir,attr"`

	VersionInfo struct {
		AppVersion string `xml:"ptKernel>appVersion"`
		DataFormat string `xml:"ptKernel>dataFormat"`
		DataBuild  string `xml:"ptKernel>dataBuild"`
	} `xml:"itdVersionInfo"`
}

func (efa *EFAProvider) postRequest(result efaResult, params url.Values) error {

	client := http.Client{}

	reqUrl, err := url.Parse(efa.BaseURL)
	if err != nil {
		return err
	}
	reqUrl.Path = path.Join(reqUrl.Path, result.endpoint())

	req, err := http.NewRequest("POST", reqUrl.String(), strings.NewReader(params.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set(
		"User-Agent",
		"GoEFA, a golang EFA client / 0.0.1 (https://github.com/michiwend/goefa)",
	)

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close() // FIXME, refer to http://golang.org/pkg/net/http/#NewRequest
	defer req.Body.Close()

	decoder := xml.NewDecoder(resp.Body)
	decoder.CharsetReader = charset.NewReader
	if err = decoder.Decode(result); err != nil {
		return err
	}

	return nil
}
