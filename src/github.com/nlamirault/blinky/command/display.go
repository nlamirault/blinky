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
	c.doDisplaySystemInformations()
	return 0
}

func (c *DisplayCommand) doDisplaySystemInformations() {
	c.UI.Info(colorstring.Color("[green] Display system informations"))
	osrelease, _ := linux.GetOSRelease()
	log.Printf("[DEBUG] OS: %s", osrelease)
	ossystem, _, _ := linux.GetKernelInformations()
	c.UI.Output(fmt.Sprintf("OS: %s %s",
		osrelease.PrettyName,
		ossystem.Architecture))
}
