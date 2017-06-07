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

// type OsFunc func() (OperatingSystem, error)

// var registeredOS = map[string](OsFunc){}

// func RegisterOperatingSystem(name string, f OsFunc) {
// 	registeredOS[name] = f
// }

// func New(name string) (OperatingSystem, error) {
// 	if os, ok := registeredOS[name]; ok {
// 		return os()
// 	}
// 	return nil, fmt.Errorf("Unsupported operating system: %s", name)
// }
