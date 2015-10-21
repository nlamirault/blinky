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

type SystemCommand struct {
	UI cli.Ui
}

func (c *SystemCommand) Help() string {
	helpText := `
Usage: blinky system [options]
	Display system informations
Options:
	--debug                       Debug mode enabled
`
	return strings.TrimSpace(helpText)
}

func (c *SystemCommand) Synopsis() string {
	return "Display system informations"
}

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
	cpuinfo, err := cpu.CPUInfo()
	if err != nil {
		c.UI.Error(colorstring.Color("[red] Erreur ") +
			fmt.Sprintf(" : %s\n", err.Error()))
		return 1
	}
	log.Printf("[DEBUG] CPU: %s", cpuinfo)
	osrelease, _ := linux.GetOSRelease()
	log.Printf("[DEBUG] OS: %s", osrelease)
	ossystem, kernel, _ := linux.GetKernelInformations()
	// platform, family, version, err := host.GetPlatformInformation()
	// if err != nil {
	// 	log.Errorf("Error: %v", err)
	// 	return
	// }
	hostInfo, err := host.HostInfo()
	if err != nil {
		c.UI.Error(colorstring.Color("[red] Erreur ") +
			fmt.Sprintf(" : %s\n", err.Error()))
		return 1
	}
	log.Printf("[DEBUG] Host: %s", hostInfo)
	vmem, err := mem.VirtualMemory()
	if err != nil {
		c.UI.Error(colorstring.Color("[red] Erreur ") +
			fmt.Sprintf(" : %s\n", err.Error()))
		return 1
	}
	// Display system informations
	c.UI.Output(colorstring.Color("[blue]OS: ") +
		fmt.Sprintf("%s %s", osrelease.PrettyName, ossystem.Architecture))
	c.UI.Output(colorstring.Color("[blue]Hostname: ") + ossystem.Hostname)
	c.UI.Output(colorstring.Color("[blue]Kernel: ") + kernel.Release)
	c.UI.Output(colorstring.Color("[blue]Memory: ") +
		fmt.Sprintf("%d/%d %d", vmem.Free, vmem.Total, vmem.UsedPercent))
	c.UI.Output(colorstring.Color("[blue]Processor: ") + cpuinfo[0].ModelName)
	// fmt.Println(chalk.Blue, "Platform:", chalk.Reset,
	// 	platform, family, version)
	c.UI.Output(colorstring.Color("[blue]Uptime: ") +
		fmt.Sprintf("%d", hostInfo.Uptime))
	return 0
}
