// Copyright (C) 2015-2017 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// +build linux

package linux

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

func (linux linuxOS) GetModel() (string, error) {
	var buffer bytes.Buffer

	productName, err := ioutil.ReadFile("/sys/devices/virtual/dmi/id/product_name")
	if err != nil {
		return "", err
	}

	productVersion, err := ioutil.ReadFile("/sys/devices/virtual/dmi/id/product_version")
	if err != nil {
		return "", err
	}

	buffer.WriteString(fmt.Sprintf("%s %s", strings.TrimSpace(string(productName)), strings.TrimSpace(string(productVersion))))
	return buffer.String(), nil
}
