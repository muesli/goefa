/*
 * Copyright (C) 2018 Christian Muehlhaeuser
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
 *   Christian Muehlhaeuser <muesli@gmail.com>
 */

package goefa

// MeansOfTransport represents the Means of Transportation
type MeansOfTransport struct {
	Name   string `xml:"name,attr"`
	Number string `xml:"symbol,attr"`

	MotType MotType `xml:"motType,attr"`
}

var motMap = map[int]string{
	-1: "Fußweg",
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

// MotType represents the Methods Of Transportation
type MotType int

func (e *MotType) String() string {
	return motMap[int(*e)]
}
