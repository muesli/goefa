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
 *      Michael Wendland <michael@michiwend.com>
 */

package goefa

// Map of supported EFA providers
var Providers = map[string]EFAProvider{

	"avv": EFAProvider{
		Name:                     "Augsburger Verkehrs- und Tarifverbund GmbH",
		BaseURL:                  "http://efa.avv-augsburg.de/avv/",
		DepartureMonitorEndpoint: "XML_DM_REQUEST",
		StopFinderEndpoint:       "XML_STOPFINDER_REQUEST",
		TripEndpoint:             "XML_TRIP_REQUEST",
	},

	"mvv": EFAProvider{
		Name:                     "Münchner Verkehrs- und Tarifverbund",
		BaseURL:                  "http://efa.mvv-muenchen.de/mvv/",
		DepartureMonitorEndpoint: "XML_DM_REQUEST",
		StopFinderEndpoint:       "XML_STOPFINDER_REQUEST",
		TripEndpoint:             "XML_TRIP_REQUEST",
	},
}
