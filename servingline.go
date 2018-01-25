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

// ServingLine holds the data for a public transportation line
type ServingLine struct {
	// ROP        int    `xml:"ROP,attr"`
	// STT        int    `xml:"displayName,attr"`
	// TTB        int    `xml:"TTB,attr"`
	// Compound   int    `xml:"compound,attr"`
	// Code       int    `xml:"code,attr"`
	// Index      string `xml:"index,attr"`
	Number     string  `xml:"number,attr"`
	Direction  string  `xml:"direction,attr"`
	DestStopID int     `xml:"destID,attr"`
	MotType    MotType `xml:"motType,attr"`
}
