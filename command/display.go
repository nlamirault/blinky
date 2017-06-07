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

package command

import (
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/mitchellh/cli"
	"github.com/mitchellh/colorstring"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"

	"github.com/nlamirault/blinky/os"
	"github.com/nlamirault/blinky/utils"
)

// DisplayCommand represents the CLI command to display logos and system
// informations
type DisplayCommand struct {
	UI cli.Ui
}

// Help defines the CLI command's help
func (c *DisplayCommand) Help() string {
	helpText := `
Usage: blinky display [options]
	Display system informations
Options:
	--debug                       Debug mode enabled
`
	return strings.TrimSpace(helpText)
}

// Synopsis defines the CLI command's synopsis
func (c *DisplayCommand) Synopsis() string {
	return "Display system informations"
}

// Run defines the CLI command
func (c *DisplayCommand) Run(args []string) int {
	var debug bool
	f := flag.NewFlagSet("display", flag.ContinueOnError)
	f.Usage = func() { c.UI.Output(c.Help()) }
	f.BoolVar(&debug, "debug", false, "Debug mode enabled")
	if err := f.Parse(args); err != nil {
		return 1
	}
	setupLogging(debug)
	return c.doDisplaySystemInformations()
}

func (c *DisplayCommand) doDisplaySystemInformations() int {
	c.UI.Output("")
	log.Printf("[DEBUG] Display system informations")

	hostInfo, err := host.Info()
	if err != nil {
		c.UI.Error(fmt.Sprintf("Error : %s", err.Error()))
		return 1
	}
	log.Printf("[DEBUG] Host: %s", hostInfo)

	cpuInfo, err := cpu.Info()
	if err != nil {
		c.UI.Error(fmt.Sprintf("Error : %s", err.Error()))
		return 1
	}
	log.Printf("[DEBUG] CPU: %s", cpuInfo)

	vmem, err := mem.VirtualMemory()
	if err != nil {
		c.UI.Error(fmt.Sprintf("Error : %s", err.Error()))
		return 1
	}

	logo, color, err := utils.GetOperatingSystemTheme(hostInfo.Platform)
	if err != nil {
		c.UI.Error(fmt.Sprintf("Error : %s", err.Error()))
		return 1
	}

	uptime, err := time.ParseDuration(fmt.Sprintf("%ds", hostInfo.Uptime))
	if err != nil {
		c.UI.Error(fmt.Sprintf("Error : %s", err.Error()))
		return 1
	}

	opSystem, err := os.NewOperatingSystem()
	if err != nil {
		c.UI.Error(fmt.Sprintf("Error : %s", err.Error()))
		return 1
	}

	details, err := os.RetrieveDetails(opSystem)
	if err != nil {
		c.UI.Error(fmt.Sprintf("Error : %s", err.Error()))
		return 1
	}

	c.UI.Output(fmt.Sprintf(
		logo,
		colorstring.Color(fmt.Sprintf("[%s]OS", color)),
		fmt.Sprintf("%s", details.Name),
		colorstring.Color(fmt.Sprintf("[%s]Model", color)),
		fmt.Sprintf("%s", details.Model),
		colorstring.Color(fmt.Sprintf("[%s]Kernel", color)),
		hostInfo.KernelVersion,
		colorstring.Color(fmt.Sprintf("[%s]Hostname", color)),
		hostInfo.Hostname,
		colorstring.Color(fmt.Sprintf("[%s]Uptime", color)),
		fmt.Sprintf("%s", uptime),
		colorstring.Color(fmt.Sprintf("[%s]Processor", color)),
		cpuInfo[0].ModelName,
		colorstring.Color(fmt.Sprintf("[%s]Mem", color)),
		fmt.Sprintf("%d MiB / %d MiB - %.0f%%", (vmem.Free/1024/1024), (vmem.Total/1024/1024), vmem.UsedPercent),
		colorstring.Color(fmt.Sprintf("[%s]Desktop", color)),
		fmt.Sprintf("%s", details.Desktop),
		colorstring.Color(fmt.Sprintf("[%s]Shell", color)),
		fmt.Sprintf("%s", details.Shell),
	))
	return 0
}
