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

package os

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const (
	productNameFile    = "/sys/devices/virtual/dmi/id/product_name"
	productVersionFile = "/sys/devices/virtual/dmi/id/product_version"
	deviceModelFile    = "/sys/firmware/devicetree/base/model"
)

func readProductFile() ([]byte, error) {
	if _, err := os.Stat(productNameFile); err == nil {
		productName, err := ioutil.ReadFile(productNameFile)
		if err != nil {
			return nil, err
		}
		return productName, nil
	} else if _, err := os.Stat(deviceModelFile); err == nil {
		productName, err := ioutil.ReadFile(deviceModelFile)
		if err != nil {
			return nil, err
		}
		return productName, nil
	}
	return []byte{}, nil
}

func readVersionFile() ([]byte, error) {
	if _, err := os.Stat(productVersionFile); err == nil {
		productVersion, err := ioutil.ReadFile(productVersionFile)
		if err != nil {
			return nil, err
		}
		return productVersion, nil
	}
	return []byte{}, nil
}

func (linux linuxOS) GetModel() (string, error) {
	productName, err := readProductFile()
	if err != nil {
		return "", err
	}

	productVersion, err := readVersionFile()
	if err != nil {
		return "", err
	}

	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("%s %s", strings.TrimSpace(string(productName)), strings.TrimSpace(string(productVersion))))
	return buffer.String(), nil
}
