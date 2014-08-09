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

type DateTimeType struct {
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

type Line struct {
	Number    string `xml:"number,attr"`
	Direction string `xml:"direction,attr"`
}

type Departure struct {
	Countdown int    `xml:"countdown,attr"`
	Platform  string `xml:"platform,attr"`

	DateTime    DateTimeType `xml:"itdDateTime"`
	ServingLine Line         `xml:"itdServingLine"`
}

type StopInfo struct {
	State string `xml:"state,attr"`

	IdfdStop struct {
		StopName  string `xml:",chardata"`
		MatchQlty int    `xml:"matchQuality,attr"`
		StopID    int    `xml:"stopID,attr"`
	} `xml:"odvNameElem"`
}

type StopResult struct {
	Stop       StopInfo    `xml:"itdDepartureMonitorRequest>itdOdv>itdOdvName"`
	Departures []Departure `xml:"itdDepartureMonitorRequest>itdDepartureList>itdDeparture"`
}
