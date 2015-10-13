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

package commands

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/ttacon/chalk"

	"github.com/nlamirault/blinky/linux"
)

var commandDisplay = cli.Command{
	Name:        "display",
	Usage:       "Display system informations",
	Description: ``,
	Action:      doDisplaySystemInformations,
	Flags: []cli.Flag{
		verboseFlag,
	},
}

func doDisplaySystemInformations(c *cli.Context) {
	log.Debugf("Display system informations")
	osrelease, _ := linux.GetOSRelease()
	log.Debugf("OS: %s", osrelease)
	ossystem, _, _ := linux.GetKernelInformations()
	fmt.Println(chalk.Blue, "OS:", chalk.Reset,
		osrelease.PrettyName,
		" ",
		ossystem.Architecture)
}
