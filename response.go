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

package goefa

import "encoding/xml"

// Response is the basic API response for all EFA requests
type Response struct {
	XMLName xml.Name `xml:"itdRequest"`

	Client     string `xml:"client,attr"`
	ClientIP   string `xml:"clientIP,attr"`
	Language   string `xml:"language,attr"`
	LengthUnit string `xml:"lengthUnit,attr"`
	Now        string `xml:"now,attr"`
	NowWD      int    `xml:"nowID,attr"`
	ServerID   string `xml:"serverID,attr"`
	SessionID  int    `xml:"sessionID,attr"`
	Version    string `xml:"version,attr"`
	VirtDir    string `xml:"virtDir,attr"`

	VersionInfo struct {
		AppVersion string `xml:"ptKernel>appVersion"`
		DataFormat string `xml:"ptKernel>dataFormat"`
		DataBuild  string `xml:"ptKernel>dataBuild"`
	} `xml:"itdVersionInfo"`
}
