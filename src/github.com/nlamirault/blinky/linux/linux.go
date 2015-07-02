// Copyright (C) 2015  Nicolas Lamirault <nicolas.lamirault@gmail.com>

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

package linux

import (
	"bufio"
	"os"
	"strings"

	log "github.com/Sirupsen/logrus"
	"golang.org/x/sys/unix"

	"github.com/nlamirault/blinky/utils"
)

// OSRelease holds the OS release facts.
type OSRelease struct {
	Name       string
	ID         string
	PrettyName string
	Version    string
	VersionID  string
}

type OSSystem struct {
	Domainname   string
	Architecture string
	Hostname     string
}

// Kernel holds the kernel facts.
type Kernel struct {
	Name    string
	Release string
	Version string
}

func GetKernelInformations() (*OSSystem, *Kernel, error) {
	var buf unix.Utsname
	err := unix.Uname(&buf)
	if err != nil {
		return nil, nil, err
	}
	ossystem := new(OSSystem)
	ossystem.Domainname = utils.CharsToString(buf.Domainname)
	ossystem.Architecture = utils.CharsToString(buf.Machine)
	ossystem.Hostname = utils.CharsToString(buf.Nodename)
	kernel := new(Kernel)
	kernel.Name = utils.CharsToString(buf.Sysname)
	kernel.Release = utils.CharsToString(buf.Release)
	kernel.Version = utils.CharsToString(buf.Version)
	return ossystem, kernel, nil
}

func GetOSRelease() (*OSRelease, error) {
	osReleaseFile, err := os.Open("/etc/os-release")
	if err != nil {
		log.Errorf("Can't open OS release file : %s", err.Error())
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
