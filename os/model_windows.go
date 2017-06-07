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

// +build windows

package os

import (
	"os/exec"
)

func (windows windowsOS) GetModel() (string, error) {
	sysctl, err := exec.LookPath("wmic")
	if err != nil {
		return "", err
	}
	out, err := exec.Command(sysctl, "computersystem get manufacturer,model").Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}