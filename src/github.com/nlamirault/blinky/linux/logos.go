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

package linux

var (
	blue  = "\033[34;1m"
	cyan  = "\033[36;1m"
	logos = map[string]string{
		"arch": "\x1b[34;1m" + `                   -\` + "\n" +
			"\x1b[34;1m" + `                  .o+\` + "\n" +
			"\x1b[34;1m" + `                 \ooo/` + "\n" +
			"\x1b[34;1m" + `                \+oooo:` + "\n" +
			"\x1b[34;1m" + `               \+oooooo:` + "\n" +
			"\x1b[34;1m" + `               -+oooooo+:` + "\n" +
			"\x1b[34;1m" + `             \/:-:++oooo+:` + "\n" +
			"\x1b[34;1m" + `            \/++++/+++++++:` + "\n" +
			"\x1b[34;1m" + `           \/++++++++++++++:` + "\n" +
			"\x1b[34;1m" + `          \/+++o` + "\x1b[36;1m" + `oooooooo` + "\x1b[34;1m" + `oooo/\` + "\n" +
			"\x1b[36;1m" + `         ` + "\x1b[34;1m" + `./` + "\x1b[36;1m" + `ooosssso++osssssso` + "\x1b[34;1m" + `+\` + "\n" +
			"\x1b[36;1m" + `        .oossssso-\\\\/ossssss+\` + "\n" +
			"\x1b[36;1m" + `       -osssssso.      :ssssssso.` + "\n" +
			"\x1b[36;1m" + `      :osssssss/        osssso+++.` + "\n" +
			"\x1b[36;1m" + `     /ossssssss/        +ssssooo/-` + "\n" +
			"\x1b[36;1m" + `   \/ossssso+/:-        -:/+osssso+-` + "\n" +
			"\x1b[36;1m" + `  \+sso+:-\                 \.-/+oso:` + "\n" +
			"\x1b[36;1m" + ` \++:.                           \-/+/` + "\n" +
			"\x1b[36;1m" + ` .\                                 \/` + "\n",
		"ubuntu": "toto" +
			"tutu",
	}
)

func GetLogo(name string) string {
	return logos[name]
}
