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

import (
	"sort"
)

var (
	red    = "\x1b[31;1m"
	green  = "\x1b[32;1m"
	yellow = "\x1b[33;1m"
	blue   = "\x1b[34;1m"
	cyan   = "\x1b[36;1m"
	white  = "\x1b[37;1m"
	logos  = map[string]string{
		"arch": cyan + `                   -\` + "\n" +
			cyan + `                  .o+\` + "\n" +
			cyan + `                 \ooo/` + "\n" +
			cyan + `                \+oooo:` + "\n" +
			cyan + `               \+oooooo:` + "\n" +
			cyan + `               -+oooooo+:` + "\n" +
			cyan + `             \/:-:++oooo+:` + "\n" +
			cyan + `            \/++++/+++++++:` + "\n" +
			cyan + `           \/++++++++++++++:` + "\n" +
			cyan + `          \/+++o` + blue + `oooooooo` + cyan + `oooo/\` + "\n" +
			blue + `         ` + cyan + `./` + blue + `ooosssso++osssssso` + cyan + `+\` + "\n" +
			blue + `        .oossssso-\\\\/ossssss+\` + "\n" +
			blue + `       -osssssso.      :ssssssso.` + "\n" +
			blue + `      :osssssss/        osssso+++.` + "\n" +
			blue + `     /ossssssss/        +ssssooo/-` + "\n" +
			blue + `   \/ossssso+/:-        -:/+osssso+-` + "\n" +
			blue + `  \+sso+:-\                 \.-/+oso:` + "\n" +
			blue + ` \++:.                           \-/+/` + "\n" +
			blue + ` .\                                 \/` + "\n",
		"debian": red + `       _,met$$$$$gg.           ` + "\n" +
			red + `    ,g$$$$$$$$$$$$$$$P.       ` + "\n" +
			red + `  ,g$$P""       """Y$$.".     ` + "\n" +
			red + ` ,$$P'              \'$$$.     ` + "\n" +
			red + `',$$P       ,ggs.     \'$$b:   ` + "\n" +
			red + `\'d$$'     ,$P"\'   ` + white + `.` + red + `    $$$    ` + "\n" +
			red + ` $$P      d$\'     ` + white + `,` + red + `    $$P    ` + "\n" +
			red + ` $$:      $$.   ` + white + `-` + red + `    ,d$$'    ` + "\n" +
			red + ` $$\;      Y$b._   _,d$P'     ` + "\n" +
			red + ` Y$$.    ` + white + `\'.` + red + `\'"Y$$$$P"'         ` + "\n" +
			red + ` \'$$b      ` + white + `"-.__              ` + "\n" +
			red + `  \'Y$$                        ` + "\n" +
			red + `   \'Y$$.                      ` + "\n" +
			red + `     \'$$b.                    ` + "\n" +
			red + `       \'Y$$b.                 ` + "\n" +
			red + `          \'"Y$b._             ` + "\n" +
			red + `              \'""""           ` + "\n",
		"ubuntu": red + `                          ./+o+-       ` + "\n" +
			white + `                  yyyyy- ` + red + `-yyyyyy+     ` + "\n" +
			white + `               ` + white + `://+//////` + red + `-yyyyyyo     ` + "\n" +
			yellow + `           .++ ` + white + `.:/++++++/-` + red + `.+sss/\     ` + "\n" +
			yellow + `         .:++o:  ` + white + ` +/++++++++/:--:/-     ` + "\n" +
			yellow + `        o:+o+:++. ` + white + `\'..\'\'\'.-/oo+++++/    ` + "\n" +
			yellow + `       .:+o:+o/.` + white + `         \'+sssoo+/   ` + "\n" +
			white + `  .++/+:` + yellow + `+oo+o:\` + white + `             /sssooo.  ` + "\n" +
			white + ` /+++//+:` + yellow + `\'oo+o` + white + `               /::--:.  ` + "\n" +
			white + ` \+/+o+++` + yellow + `\'o++o` + red + `               ++////.  ` + "\n" +
			white + `  .++.o+` + yellow + ` ++oo+:\'` + red + `             /dddhhh.  ` + "\n" +
			yellow + `       .+.o+oo:.` + red + `          \'oddhhhh+   ` + "\n" +
			yellow + `        \+.++o+o\'` + red + `\'-\'\'\'\'.:ohdhhhhh+    ` + "\n" +
			yellow + `         \':o+++ ` + red + `\'ohhhhhhhhyo++os:     ` + "\n" +
			yellow + `           .o:` + red + `\'.syhhhhhhh/` + yellow + `.oo++o\'     ` + "\n" +
			red + `               /osyyyyyyo` + yellow + `++ooo+++/    ` + "\n" +
			red + `                   \'\'\'\'\' ` + yellow + `+oo+++o\:    ` + "\n" +
			yellow + `                          \'oo++.      ` + "\n",
		"mint": green + ` MMMMMMMMMMMMMMMMMMMMMMMMMmds+.       ` + "\n" +
			green + `MMm----::-://////////////oymNMd+'    ` + "\n" +
			green + `MMd      ` + white + `/++                ` + green + `-sNMd:   ` + "\n" +
			green + `MMNso/\'  ` + white + `dMM    \'.::-. .-::.\' ` + green + `.hMN:  ` + "\n" +
			green + `ddddMMh  ` + white + `dMM   :hNMNMNhNMNMNh: ` + green + `\'NMm  ` + "\n" +
			green + `    NMm  ` + white + `dMM  .NMN/-+MMM+-/NMN\' ` + green + `dMM  ` + "\n" +
			green + `    NMm  ` + white + `dMM  -MMm  \'MMM   dMM. ` + green + `dMM  ` + "\n" +
			green + `    NMm  ` + white + `dMM  -MMm  \'MMM   dMM. ` + green + `dMM  ` + "\n" +
			green + `    NMm  ` + white + `dMM  .mmd  \'mmm   yMM. ` + green + `dMM  ` + "\n" +
			green + `    NMm  ` + white + `dMM\'  ..\'   ...   ydm. ` + green + `dMM  ` + "\n" +
			green + `    hMM- ` + white + `+MMd/-------...-:sdds  ` + green + `dMM  ` + "\n" +
			green + `    -NMm- ` + white + `:hNMNNNmdddddddddy/\'  ` + green + `dMM  ` + "\n" +
			green + `     -dMNs-` + white + `\'\'-::::-------.\'\'    ` + green + `dMM  ` + "\n" +
			green + `      \'/dMNmy+/:-------------:/yMMM  ` + "\n" +
			green + `         ./ydNMMMMMMMMMMMMMMMMMMMMM  ` + "\n" +
			green + `            \.MMMMMMMMMMMMMMMMMMM    ` + "\n",
	}
)

func GetLinuxDistributions() []string {
	keys := make([]string, 0, len(logos))
	for key := range logos {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

func GetLogo(name string) string {
	return logos[name]
}
