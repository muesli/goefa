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

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

// ProviderFromJson parses providers.json. If a json object matches short_name
// a pointer to the corresponding EFAProvider is returned.
func ProviderFromJson(short_name string) (*EFAProvider, error) {

	content, err := ioutil.ReadFile("providers.json")

	if err != nil {
		return nil, err
	}

	var providers map[string]*EFAProvider
	json.Unmarshal(content, &providers)

	provider, contains := providers[short_name]

	if contains == false {
		return nil, errors.New("Provider '" + short_name + "' not found.")
	}

	return provider, nil
}
