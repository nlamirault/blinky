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
	"bufio"
	"os"
	"strings"
)

// OSRelease holds the OS release facts.
type OSRelease struct {
	Name       string
	ID         string
	PrettyName string
	Version    string
	VersionID  string
}

// GetOSRelease extract informations for the current operating system
func GetOSRelease() (*OSRelease, error) {
	osReleaseFile, err := os.Open("/etc/os-release")
	if err != nil {
		//log.Printf("[ERROR] Can't open OS release file : %s", err.Error())
		return nil, err
	}
	defer osReleaseFile.Close()
	osrelease := new(OSRelease)

	scanner := bufio.NewScanner(osReleaseFile)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			columns := strings.Split(line, "=")
			key := columns[0]
			value := strings.Trim(strings.TrimSpace(columns[1]), `"`)
			switch key {
			case "NAME":
				osrelease.Name = value
			case "ID":
				osrelease.ID = value
			case "PRETTY_NAME":
				osrelease.PrettyName = value
			case "VERSION":
				osrelease.Version = value
			case "VERSION_ID":
				osrelease.VersionID = value
			}
		}
	}
	return osrelease, nil
}
