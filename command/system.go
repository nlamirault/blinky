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

	"github.com/mitchellh/cli"
	"github.com/mitchellh/colorstring"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"

	"github.com/nlamirault/blinky/os/linux"
)

// SystemCommand represents the CLI command which display operation system
// informations
type SystemCommand struct {
	UI cli.Ui
}

// Help defines the CLI command's help
func (c *SystemCommand) Help() string {
	helpText := `
Usage: blinky system [options]
	Display system informations
Options:
	--debug                       Debug mode enabled
`
	return strings.TrimSpace(helpText)
}

// Synopsis defines the CLI command's synopsis
func (c *SystemCommand) Synopsis() string {
	return "Display system informations"
}

// Run defines the CLI command
func (c *SystemCommand) Run(args []string) int {
	var debug bool
	f := flag.NewFlagSet("system", flag.ContinueOnError)
	f.Usage = func() { c.UI.Output(c.Help()) }
	f.BoolVar(&debug, "debug", false, "Debug mode enabled")
	if err := f.Parse(args); err != nil {
		return 1
	}
	setupLogging(debug)
	return c.doSystemInformations()
}

func (c *SystemCommand) doSystemInformations() int {
	log.Printf("[DEBUG] Retrieve system informations")
	osrelease, err := linux.GetOSRelease()
	if err != nil {
		c.UI.Error(fmt.Sprintf("Error : %s", err.Error()))
		return 1
	}
	log.Printf("[DEBUG] OS: %s", osrelease)
	cpuinfo, err := cpu.Info()
	if err != nil {
		c.UI.Error(fmt.Sprintf("Error : %s", err.Error()))
		return 1
	}
	log.Printf("[DEBUG] CPU: %s", cpuinfo)

	hostInfo, err := host.Info()
	if err != nil {
		c.UI.Error(fmt.Sprintf("Error : %s", err.Error()))
		return 1
	}
	log.Printf("[DEBUG] Host: %s", hostInfo)
	vmem, err := mem.VirtualMemory()
	if err != nil {
		c.UI.Error(fmt.Sprintf("Error : %s", err.Error()))
		return 1
	}
	// Display system informations
	c.UI.Output(colorstring.Color("[blue]OS: ") +
		fmt.Sprintf("%s", hostInfo.PlatformFamily))
	c.UI.Output(colorstring.Color("[blue]Hostname: ") + hostInfo.Hostname)
	c.UI.Output(colorstring.Color("[blue]Kernel: ") + hostInfo.KernelVersion)
	c.UI.Output(colorstring.Color("[blue]Memory: ") +
		fmt.Sprintf("%d/%d %d", vmem.Free, vmem.Total, vmem.UsedPercent))
	c.UI.Output(colorstring.Color("[blue]Processor: ") + cpuinfo[0].ModelName)
	// fmt.Println(chalk.Blue, "Platform:", chalk.Reset,
	// 	platform, family, version)
	c.UI.Output(colorstring.Color("[blue]Uptime: ") +
		fmt.Sprintf("%d", hostInfo.Uptime))
	return 0
}
