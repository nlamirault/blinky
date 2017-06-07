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
	"bufio"
	"fmt"
	"os"
	"strings"
)

type linuxOS struct{}

func NewOperatingSystem() (OperatingSystem, error) {
	return linuxOS{}, nil
}

func (linux linuxOS) GetShell() (string, error) {
	return os.Getenv("SHELL"), nil
}

func (linux linuxOS) GetName() (string, error) {
	osReleaseFile, err := os.Open("/etc/os-release")
	if err != nil {
		return "", err
	}
	defer osReleaseFile.Close()
	scanner := bufio.NewScanner(osReleaseFile)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			columns := strings.Split(line, "=")
			key := columns[0]
			value := strings.Trim(strings.TrimSpace(columns[1]), `"`)
			switch key {
			case "PRETTY_NAME":
				return value, nil
			}
		}
	}
	return "", fmt.Errorf("Can't find Linux distribution name")
}
