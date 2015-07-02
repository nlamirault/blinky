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
	//"fmt"
	//"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	//"github.com/ttacon/chalk"
)

var commandSystemInfos = cli.Command{
	Name:        "infos",
	Usage:       "System informations",
	Description: ``,
	Action:      doSystemInformations,
	Flags: []cli.Flag{
		verboseFlag,
	},
}

func doSystemInformations(c *cli.Context) {
	log.Debugf("Retrieve system informations")
	// date := linux.GetDate()
	// log.Debugf("Date: %s", date)
	// osrelease, _ := linux.GetOSRelease()
	// log.Debugf("OS: %s", osrelease)
	// ossystem, kernel, _ := linux.GetKernelInformations()
	// log.Debugf("System: %s", ossystem)
	// log.Debugf("Kernel: %s", kernel)
	// memory, uptime, _ := linux.GetLoadAndMemory()
	// log.Debugf("Memory: %s", memory)
	// log.Debugf("Uptime: %s", uptime)

	// fmt.Println(chalk.Blue, "OS:", chalk.Reset,
	// 	osrelease.PrettyName,
	// 	" ",
	// 	ossystem.Architecture)
	// fmt.Println(chalk.Blue, "Hostname:", chalk.Reset,
	// 	ossystem.Hostname)
	// fmt.Println(chalk.Blue, "Kernel:", chalk.Reset,
	// 	kernel.Release)
	// fmt.Println(chalk.Blue, "Memory:", chalk.Reset,
	// 	memory.Free, "/", memory.Total)
}
