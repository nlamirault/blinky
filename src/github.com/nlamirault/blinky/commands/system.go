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
	//"os"

	log "github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/ttacon/chalk"

	"github.com/nlamirault/blinky/linux"
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
	cpuinfo, err := cpu.CPUInfo()
	if err != nil {
		log.Errorf("Error: %v", err)
		return
	}
	log.Debugf("CPU: %s", cpuinfo)
	osrelease, _ := linux.GetOSRelease()
	log.Debugf("OS: %s", osrelease)
	ossystem, kernel, _ := linux.GetKernelInformations()
	// platform, family, version, err := host.GetPlatformInformation()
	// if err != nil {
	// 	log.Errorf("Error: %v", err)
	// 	return
	// }
	hostInfo, err := host.HostInfo()
	if err != nil {
		log.Errorf("Error: %v", err)
		return
	}
	log.Debugf("Host: %s", hostInfo)
	vmem, err := mem.VirtualMemory()
	if err != nil {
		log.Errorf("Error: %v", err)
		return
	}
	// Display system informations
	fmt.Println(chalk.Blue, "OS:", chalk.Reset,
		osrelease.PrettyName,
		" ",
		ossystem.Architecture)
	fmt.Println(chalk.Blue, "Hostname:", chalk.Reset,
		ossystem.Hostname)
	fmt.Println(chalk.Blue, "Kernel:", chalk.Reset,
		kernel.Release)
	fmt.Println(chalk.Blue, "Memory:", chalk.Reset,
		vmem.Free, "/", vmem.Total, " ", vmem.UsedPercent)
	fmt.Println(chalk.Blue, "Processor:", chalk.Reset,
		cpuinfo[0].ModelName)
	// fmt.Println(chalk.Blue, "Platform:", chalk.Reset,
	// 	platform, family, version)
	fmt.Println(chalk.Blue, "Uptime:", chalk.Reset,
		hostInfo.Uptime)
}
