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

// Provider represents a public transport company that provides access to
// its EFA instance. Use providers.json to store a list of known providers.
type Provider struct {
	BaseURL        string
	EnableRealtime bool
}

// NewProvider returns a new Provider with custom settings
func NewProvider(baseurl string, realtime bool) *Provider {
	return &Provider{
		BaseURL:        baseurl,
		EnableRealtime: realtime,
	}
}
