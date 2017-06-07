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

// +build darwin

package os

import (
	"fmt"
	"os"
	"regexp"

	"github.com/nlamirault/blinky/utils"
)

var (
	tiger        = regexp.MustCompile("10.4.[0-9]+")
	leopard      = regexp.MustCompile("10.5.[0-9]+")
	snowleopard  = regexp.MustCompile("10.6.[0-9]+")
	lion         = regexp.MustCompile("10.7.[0-9]+")
	mountainlion = regexp.MustCompile("10.8.[0-9]+")
	mavericks    = regexp.MustCompile("10.9.[0-9]+")
	yosemite     = regexp.MustCompile("10.10.[0-9]+")
	elcapitan    = regexp.MustCompile("10.11.[0-9]+")
	sierra       = regexp.MustCompile("10.12.[0-9]+")
)

type darwinOS struct{}

func NewOperatingSystem() (OperatingSystem, error) {
	return darwinOS{}, nil
}

func (darwin darwinOS) GetShell() (string, error) {
	return os.Getenv("SHELL"), nil
}

func (darwin darwinOS) GetName() (string, error) {
	osxVersion, err := utils.ExecCommand("sw_vers", "-productVersion")
	if err != nil {
		return "", err
	}
	osxBuild, err := utils.ExecCommand("sw_vers", "-buildVersion")
	if err != nil {
		return "", err
	}
	codename := "Mac OS"
	switch {
	case tiger.MatchString(osxVersion):
		codename = "Mac OS X Tiger"
	case leopard.MatchString(osxVersion):
		codename = "Mac OS X Leopard"
	case snowleopard.MatchString(osxVersion):
		codename = "Mac OS X Snow Leopard"
	case lion.MatchString(osxVersion):
		codename = "Mac OS X Lion"
	case mountainlion.MatchString(osxVersion):
		codename = "Mac OS X Mountain Lion"
	case mavericks.MatchString(osxVersion):
		codename = "Mac OS X Mavericks"
	case yosemite.MatchString(osxVersion):
		codename = "Mac OS X Yosemite"
	case elcapitan.MatchString(osxVersion):
		codename = "Mac OS X El Capitan"
	case sierra.MatchString(osxVersion):
		codename = "macOS Sierra"
	}
	return fmt.Sprintf("%s %s %s", codename, osxVersion, osxBuild), nil
}
