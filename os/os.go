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

package os

// import (
// 	"fmt"
// )

type OperatingSystem interface {

	// GetModel return the computer name
	GetModel() (string, error)

	// GetDesktop return the name of the window manager or desktop manager
	GetDesktop() (string, error)

	// GetShell return the Shell used
	GetShell() (string, error)

	GetName() (string, error)
}

// Details define some informatinos about the operating system
type Details struct {
	Model   string
	Desktop string
	Name    string
	Shell   string
}

func RetrieveDetails(system OperatingSystem) (*Details, error) {
	model, err := system.GetModel()
	if err != nil {
		return nil, err
	}

	desktopName, err := system.GetDesktop()
	if err != nil {
		return nil, err
	}

	osName, err := system.GetName()
	if err != nil {
		return nil, err
	}

	shellName, err := system.GetShell()
	if err != nil {
		return nil, err
	}
	return &Details{
		Model:   model,
		Desktop: desktopName,
		Name:    osName,
		Shell:   shellName,
	}, nil
}
