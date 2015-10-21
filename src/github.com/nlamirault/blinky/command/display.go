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

package command

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/mitchellh/cli"
	"github.com/mitchellh/colorstring"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"

	"github.com/nlamirault/blinky/linux"
)

type DisplayCommand struct {
	UI cli.Ui
}

func (c *DisplayCommand) Help() string {
	helpText := `
Usage: blinky display [options]
	Display system informations
Options:
	--debug                       Debug mode enabled
`
	return strings.TrimSpace(helpText)
}

func (c *DisplayCommand) Synopsis() string {
	return "Display system informations"
}

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
	c.UI.Info("Display system informations")
	osrelease, err := linux.GetOSRelease()
	if err != nil {
		c.UI.Error(fmt.Sprintf("Error : %s", err.Error()))
		return 1
	}
	log.Printf("[DEBUG] OS: %s", osrelease)
	logo := linux.GetLogoFormat(osrelease.ID)
	ossystem, kernel, err := linux.GetKernelInformations()
	if err != nil {
		c.UI.Error(fmt.Sprintf("Error : %s", err.Error()))
		return 1
	}
	hostInfo, err := host.HostInfo()
	if err != nil {
		c.UI.Error(fmt.Sprintf("Error : %s", err.Error()))
		return 1
	}
	cpuinfo, err := cpu.CPUInfo()
	if err != nil {
		c.UI.Error(fmt.Sprintf("Error : %s", err.Error()))
		return 1
	}
	vmem, err := mem.VirtualMemory()
	if err != nil {
		c.UI.Error(fmt.Sprintf("Error : %s", err.Error()))
		return 1
	}

	c.UI.Output(fmt.Sprintf(
		logo,
		colorstring.Color("[blue]OS"),
		fmt.Sprintf("%s %s",
			osrelease.PrettyName, ossystem.Architecture),
		colorstring.Color("[blue]Kernel"),
		kernel.Release,
		colorstring.Color("[blue]Hostname"),
		ossystem.Hostname,
		colorstring.Color("[blue]Uptime"),
		fmt.Sprintf("%d", hostInfo.Uptime),
		colorstring.Color("[blue]Processor"),
		cpuinfo[0].ModelName,
		colorstring.Color("[blue]Mem"),
		fmt.Sprintf("%d/%d %3d", vmem.Free, vmem.Total, vmem.UsedPercent),
	))
	return 0
}
