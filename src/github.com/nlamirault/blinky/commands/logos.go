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

	"github.com/nlamirault/blinky/linux"
)

var commandLogoList = cli.Command{
	Name:        "list",
	Usage:       "List linux distributions",
	Description: ``,
	Action:      doLogoList,
	Flags: []cli.Flag{
		verboseFlag,
	},
}

var commandLogoDist = cli.Command{
	Name:        "dist",
	Usage:       "Show distribution logo",
	Description: ``,
	Action:      doLogoDist,
	Flags: []cli.Flag{
		verboseFlag,
		cli.StringFlag{
			Name:  "name",
			Usage: "Linux distribution name",
		},
	},
}

func doLogoList(c *cli.Context) {
	log.Debugf("List logos")
}

func doLogoDist(c *cli.Context) {
	log.Debugf("Distribution logo: " + c.String("name"))
	fmt.Println(linux.GetLogo(c.String("name")))

}
