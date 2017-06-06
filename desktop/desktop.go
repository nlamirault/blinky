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

// +build linux

package desktop

import (
	"fmt"

	"github.com/nlamirault/blinky/os/linux"
)

var (
	desktops = map[string]string{
		"gnome-session": "GNOME",
		"ksmserver":     "KDE",
		"mate-session":  "MATE",
		"xfce4-session": "XFCE",
		"lxsession":     "LXDE",
		"cinnamon":      "CINNAMON",
		"darwin":        "AQUA",
		"windows":       "AERO",
	}

	windowmanagers = map[string]string{
		"awesome":       "Awesome",
		"beryl":         "Beryl",
		"blackbox":      "Blackbox",
		"bspwm":         "bspwm",
		"dwm":           "DWM",
		"enlightenment": "Enlightenment",
		"fluxbox":       "Fluxbox",
		"fvwm":          "FVWM",
		"herbstluftwm":  "herbstluftwm",
		"i3":            "i3",
		"icewm":         "IceWM",
		"kwin":          "KWin",
		"metacity":      "Metacity",
		"musca":         "Musca",
		"openbox":       "Openbox",
		"pekwm":         "PekWM",
		"ratpoison":     "ratpoison",
		"scrotwm":       "ScrotWM",
		"subtle":        "subtle",
		"monsterwm":     "MonsterWM",
		"wmaker":        "Window Maker",
		"wmfs":          "Wmfs",
		"wmii":          "wmii",
		"xfwm4":         "Xfwm",
		"emerald":       "Emerald",
		"compiz":        "Compiz",
		"xmonad":        "xmonad",
		"qtile":         "QTile",
		"wingo":         "Wingo",
	}
)

func GetName(osName string) (string, error) {
	switch osName {
	case "darwin":
		return desktops[osName], nil
	case "windows":
		return desktops[osName], nil
	case "linux":
		return linux.GetDesktop()
	default:
		return "", fmt.Errorf("Operating system not supported")
	}
}
