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

package goefa

var MOTMap = map[int]string{
	0:  "Zug",
	1:  "S-Bahn",
	2:  "U-Bahn",
	3:  "Stadtbahn",
	4:  "Straßen-/Trambahn",
	5:  "Stadtbus",
	6:  "Regionalbus",
	7:  "Schnellbus",
	8:  "Seil-/Zahnradbahn",
	9:  "Schiff",
	10: "AST/Rufbus",
	11: "Sonstige",
}

type EFAMotType int

func (e *EFAMotType) String() string {
	return MOTMap[int(*e)]
}

type EFAServingLine struct {
	ROP       int    `xml:"ROP,attr"`
	STT       int    `xml:"displayName,attr"`
	TTB       int    `xml:"TTB,attr"`
	Code      int    `xml:"code,attr"`
	Compound  int    `xml:"compound,attr"`
	DestID    int    `xml:"destID,attr"`
	Direction string `xml:"direction,attr"`
	Index     string `xml:"index,attr"`
	Number    string `xml:"number,attr"`

	MotType EFAMotType `xml:"motType,attr"`

	DestStopID int `xml:"destID"` //FIXME assign EFAStop
}
