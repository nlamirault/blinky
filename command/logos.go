// Copyright (C) 2015-2017  Nicolas Lamirault <nicolas.lamirault@gmail.com>

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

	"github.com/nlamirault/blinky/utils"
)

// LogoCommand defines the command which manage operating systems logos.
type LogoCommand struct {
	UI cli.Ui
}

// Help defines the CLI command's help
func (c *LogoCommand) Help() string {
	helpText := `
Usage: blinky logo [options] action
	Operating systems logos
Options:
	--debug                       Debug mode enabled
        --name                        Operating system name

Action:
        list      : List all operating systems
        display   : Display operating system logo
`
	return strings.TrimSpace(helpText)
}

// Synopsis defines the CLI command's synopsis
func (c *LogoCommand) Synopsis() string {
	return "Operating systems logos"
}

// Run defines the CLI command
func (c *LogoCommand) Run(args []string) int {
	var debug bool
	var name string
	f := flag.NewFlagSet("display", flag.ContinueOnError)
	f.Usage = func() { c.UI.Output(c.Help()) }
	f.BoolVar(&debug, "debug", false, "Debug mode enabled")
	f.StringVar(&name, "name", "", "Operating system name")
	if err := f.Parse(args); err != nil {
		return 1
	}
	action := f.Args()
	//fmt.Printf("Args : %s %s\n", action, name)
	if len(action) != 1 {
		errorMessage(
			c.UI,
			"At least one action to logo must be specified.",
			c.Help())
		return 1
	}
	setupLogging(debug)
	if action[0] == "list" {
		c.doLogoList()
		return 0
	} else if action[0] == "display" {
		if len(name) > 0 {
			c.doLogoDist(name)
			return 0
		}
		errorMessage(c.UI, "Please specify name", c.Help())
		return 1
	}
	return 0
}

func (c *LogoCommand) doLogoList() {
	c.UI.Info(colorstring.Color("[green] List logos"))
	dists := utils.GetOperatingSystems()
	log.Printf("[DEBUG] Linux distributions: %s", dists)
	for _, dist := range dists {
		c.UI.Output(fmt.Sprintf(" - %s", dist))
	}
}

func (c *LogoCommand) doLogoDist(name string) {
	log.Printf("[DEBUG] Distribution logo: " + name)
	c.UI.Output(utils.GetLogo(name))
}
