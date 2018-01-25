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

import (
	"encoding/xml"
	"time"
)

// EFATime implements UnmarshalXML to support unmarshalling EFAs XML DateTime
// type directly into a time.Time compatible type
type EFATime struct {
	time.Time
}

// UnmarshalXML is a custom XML decoding method for EFATime
func (t *EFATime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type efaDateTime struct {
		Date struct {
			Day   int `xml:"day,attr"`
			Month int `xml:"month,attr"`
			Year  int `xml:"year,attr"`
		} `xml:"itdDate"`

		Time struct {
			Hour   int `xml:"hour,attr"`
			Minute int `xml:"minute,attr"`
		} `xml:"itdTime"`
	}

	var content efaDateTime
	if err := d.DecodeElement(&content, &start); err != nil {
		return err
	}

	loc, err := time.LoadLocation("Local")
	if err != nil {
		return err
	}

	tmp := time.Date(content.Date.Year,
		time.Month(content.Date.Month),
		content.Date.Day,
		content.Time.Hour,
		content.Time.Minute,
		0,
		0,
		loc)

	t.Time = tmp

	return nil
}
