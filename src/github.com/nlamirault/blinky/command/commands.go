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
	//"log"
	"os"

	"github.com/mitchellh/cli"

	"github.com/nlamirault/blinky/logging"
)

// Commands is the mapping of all the available Terraform commands.
var (
	Commands map[string]cli.CommandFactory
	UI       cli.Ui
)

type Meta struct {
	UI cli.Ui
}

func init() {
	UI = &cli.ColoredUi{
		Ui: &cli.BasicUi{
			Writer:      os.Stdout,
			Reader:      os.Stdin,
			ErrorWriter: os.Stderr,
		},
		OutputColor: cli.UiColorNone,
		InfoColor:   cli.UiColorGreen,
		ErrorColor:  cli.UiColorRed,
	}

	Commands = map[string]cli.CommandFactory{
		"system": func() (cli.Command, error) {
			return &SystemCommand{
				UI: UI,
			}, nil
		},
		"logo": func() (cli.Command, error) {
			return &LogoCommand{
				UI: UI,
			}, nil
		},
		"display": func() (cli.Command, error) {
			return &DisplayCommand{
				UI: UI,
			}, nil
		},
	}
}

func errorMessage(ui cli.Ui, msg string, help string) {
	ui.Error(msg)
	ui.Error("")
	ui.Error(help)
}

func setupLogging(debug bool) {
	if debug {
		logging.SetLogging("DEBUG")
	} else {
		logging.SetLogging("INFO")
	}
}
