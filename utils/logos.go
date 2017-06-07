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

package utils

import (
	"fmt"
	"sort"
)

var (
	orange = "\x1b[33m"
	red    = "\x1b[31;1m"
	green  = "\x1b[32;1m"
	yellow = "\x1b[33;1m"
	blue   = "\x1b[34;1m"
	purple = "\x1b[35;1m"
	cyan   = "\x1b[36;1m"
	white  = "\x1b[37;1m"
	logos  = map[string]string{
		"arch": cyan + `                   -` + "\n" +
			cyan + `                  o+` + "\n" +
			cyan + `                 ooo` + "\n" +
			cyan + `                +oooo:` + "\n" +
			cyan + `               +oooooo:` + "\n" +
			cyan + `              --+oooooo+:` + "\n" +
			cyan + `             /:-:++oooo+:` + "\n" +
			cyan + `            /++++/+++++++:` + "\n" +
			cyan + `           /++++++++++++++:` + "\n" +
			cyan + `          /+++o` + blue + `oooooooo` + cyan + `oooo/` + "\n" +
			blue + `         ` + cyan + `./` + blue + `ooosssso++osssssso` + cyan + `+` + "\n" +
			blue + `        .oossssso-\\\\/ossssss+` + "\n" +
			blue + `       -osssssso.      :ssssssso.` + "\n" +
			blue + `      :osssssss/        osssso+++.` + "\n" +
			blue + `     /ossssssss/        +ssssooo/-` + "\n" +
			blue + `    /ossssso+/:-        -:/+osssso+-` + "\n" +
			blue + `   +sso+:-\                 \.-/+oso:` + "\n" +
			blue + `  ++:.                           \-/+/` + "\n" +
			blue + ` .\                                 \/` + "\n",
		"debian": red + `       _,met$$$$$gg.           ` + "\n" +
			red + `    ,g$$$$$$$$$$$$$$$P.       ` + "\n" +
			red + `  ,g$$P""       """Y$$.".     ` + "\n" +
			red + ` ,$$P'              '$$$.     ` + "\n" +
			red + `',$$P       ,ggs.     '$$b:   ` + "\n" +
			red + `'d$$'     ,$P"'   ` + white + `.` + red + `    $$$    ` + "\n" +
			red + ` $$P      d$'     ` + white + `,` + red + `    $$P    ` + "\n" +
			red + ` $$:      $$.   ` + white + `-` + red + `    ,d$$'    ` + "\n" +
			red + ` $$\;      Y$b._   _,d$P'     ` + "\n" +
			red + ` Y$$.    ` + white + `'.` + red + `'"Y$$$$P"'         ` + "\n" +
			red + ` '$$b      ` + white + `"-.__              ` + "\n" +
			red + `  'Y$$                        ` + "\n" +
			red + `   'Y$$.                      ` + "\n" +
			red + `     '$$b.                    ` + "\n" +
			red + `       'Y$$b.                 ` + "\n" +
			red + `          '"Y$b._             ` + "\n" +
			red + `              '""""           ` + "\n",
		"ubuntu": red + `                          ./+o+-       ` + "\n" +
			white + `                  yyyyy- ` + red + `-yyyyyy+     ` + "\n" +
			white + `               ` + white + `://+//////` + red + `-yyyyyyo     ` + "\n" +
			yellow + `           .++ ` + white + `.:/++++++/-` + red + `.+sss/\     ` + "\n" +
			yellow + `         .:++o:  ` + white + ` +/++++++++/:--:/-     ` + "\n" +
			yellow + `        o:+o+:++. ` + white + `'..'''.-/oo+++++/    ` + "\n" +
			yellow + `       .:+o:+o/.` + white + `         '+sssoo+/   ` + "\n" +
			white + `  .++/+:` + yellow + `+oo+o:` + white + `             /sssooo.  ` + "\n" +
			white + ` /+++//+:` + yellow + `'oo+o` + white + `               /::--:.  ` + "\n" +
			white + ` \+/+o+++` + yellow + `'o++o` + red + `               ++////.  ` + "\n" +
			white + `  .++.o+` + yellow + ` ++oo+:'` + red + `             /dddhhh.  ` + "\n" +
			yellow + `       .+.o+oo:.` + red + `          \'oddhhhh+   ` + "\n" +
			yellow + `        \+.++o+o\'` + red + `'-''''.:ohdhhhhh+    ` + "\n" +
			yellow + `         ':o+++ ` + red + `'ohhhhhhhhyo++os:     ` + "\n" +
			yellow + `           .o:` + red + `'.syhhhhhhh/` + yellow + `.oo++o'     ` + "\n" +
			red + `               /osyyyyyyo` + yellow + `++ooo+++/    ` + "\n" +
			red + `                   ''''' ` + yellow + `+oo+++o\:    ` + "\n" +
			yellow + `                          'oo++.      ` + "\n",
		"mint": green + `MMMMMMMMMMMMMMMMMMMMMMMMMmds+.       ` + "\n" +
			green + `MMm----::-://////////////oymNMd+'    ` + "\n" +
			green + `MMd      ` + white + `/++                ` + green + `-sNMd:   ` + "\n" +
			green + `MMNso/'  ` + white + `dMM    '.::-. .-::.' ` + green + `.hMN:  ` + "\n" +
			green + `ddddMMh  ` + white + `dMM   :hNMNMNhNMNMNh: ` + green + `'NMm  ` + "\n" +
			green + `    NMm  ` + white + `dMM  .NMN/-+MMM+-/NMN' ` + green + `dMM  ` + "\n" +
			green + `    NMm  ` + white + `dMM  -MMm  'MMM   dMM. ` + green + `dMM  ` + "\n" +
			green + `    NMm  ` + white + `dMM  -MMm  'MMM   dMM. ` + green + `dMM  ` + "\n" +
			green + `    NMm  ` + white + `dMM  .mmd  'mmm   yMM. ` + green + `dMM  ` + "\n" +
			green + `    NMm  ` + white + `dMM'  ..'   ...   ydm. ` + green + `dMM  ` + "\n" +
			green + `    hMM- ` + white + `+MMd/-------...-:sdds  ` + green + `dMM  ` + "\n" +
			green + `    -NMm- ` + white + `:hNMNNNmdddddddddy/'  ` + green + `dMM  ` + "\n" +
			green + `     -dMNs-` + white + `''-::::-------.''    ` + green + `dMM  ` + "\n" +
			green + `      '/dMNmy+/:-------------:/yMMM  ` + "\n" +
			green + `         ./ydNMMMMMMMMMMMMMMMMMMMMM  ` + "\n" +
			green + `            \.MMMMMMMMMMMMMMMMMMM    ` + "\n",
		"centos": yellow +
			`                 .PLTJ.                  ` + "\n" +
			`                <><><><>                 ` + "\n" +
			green +
			`      KKSSV' 4KKK  ` + yellow + `LJ ` + red + ` KKKL.'VSSKK ` + "\n" +
			green + `      KKV' 4KKKKK ` + yellow + ` LJ` + red + `  KKKKAL 'VKK` + "\n" +
			green + `      V' ' 'VKKKK ` + yellow + ` LJ ` + red + ` KKKKV' ' 'V` + "\n" +
			green + `      .4MA.' 'VKK ` + yellow + ` LJ ` + red + ` KKV' '.4Mb.` + "\n" +
			red + `    . ` + green + `KKKKKA.' VK ` + yellow + ` LJ ` + red + ` V' '.4KKKKK` + blue + `.` + "\n" +
			red + `  .4D ` + green + `KKKKKKKA.'' ` + yellow + ` LJ ` + red + ` ''.4KKKKKKK ` + blue + `FA.` + "\n" +
			red + ` <QDD ++++++++++++  ` + blue + `  ++++++++++++ GFD>` + "\n" +
			red + `  'VD ` + blue + `KKKKKKKK'.. ` + green + ` LJ ` + yellow + ` ..'KKKKKKKK ` + blue + `FV` + "\n" +
			red + `    ' ` + blue + `VKKKKK'. .4 ` + green + ` LJ  ` + yellow + `K. .'KKKKKV` + blue + `'` + "\n" +
			blue + `      'VK'. .4KK   ` + green + `LJ ` + yellow + ` KKA. .'KV'` + "\n" +
			blue + `      A. . .4KKKK  ` + green + `LJ  ` + yellow + `KKKKA. . .4` + "\n" +
			blue + `      KKA. 'KKKKK  ` + green + `LJ  ` + yellow + `KKKKK' .4KK` + "\n" +
			blue + `      KKSSA. VKKK  ` + green + `LJ  ` + yellow + `KKKV .4SSKK` + "\n" +
			green + `                <><><><>` + "\n" +
			green + `                 'MKKM'` + "\n",
		"osx": green + `                  -/+:.          ` + "\n" +
			green + `                 :++++.         ` + "\n" +
			green + `                /+++/.          ` + "\n" +
			green + `        .:-::- .+/:-''.::-      ` + "\n" +
			green + `     .:/++++++/::::/++++++/:'   ` + "\n" +
			yellow + `  .:///////////////////////:'  ` + "\n" +
			yellow + `  ////////////////////////'    ` + "\n" +
			orange + ` -+++++++++++++++++++++++'     ` + "\n" +
			orange + ` /++++++++++++++++++++++/      ` + "\n" +
			red + ` /sssssssssssssssssssssss.     ` + "\n" +
			red + ` :ssssssssssssssssssssssss-    ` + "\n" +
			purple + `  osssssssssssssssssssssssso/' ` + "\n" +
			purple + `  'syyyyyyyyyyyyyyyyyyyyyyyy+' ` + "\n" +
			blue + `   'ossssssssssssssssssssss/   ` + "\n" +
			blue + `     :ooooooooooooooooooo+.    ` + "\n" +
			blue + `      ':+oo+/:-..-:/+o+/-      ` + "\n",
		"windows": red + `        ,.=:!!t3Z3z.,                 ` + "\n" +
			red + `       :tt:::tt333EE3                ` + "\n" +
			red + `       Et:::ztt33EEEL` + green + ` @Ee.,      .., ` + "\n" +
			red + `      ;tt:::tt333EE7` + green + ` ;EEEEEEttttt33# ` + "\n" +
			red + `     :Et:::zt333EEQ.` + green + ` $EEEEEttttt33QL ` + "\n" +
			red + `     it::::tt333EEF` + green + ` @EEEEEEttttt33F  ` + "\n" +
			red + `    ;3=*^'''"*4EEV` + green + ` :EEEEEEttttt33@.  ` + "\n" +
			blue + `    ,.=::::!t=., ` + red + `'` + green + ` @EEEEEEtttz33QF   ` + "\n" +
			blue + `   ;::::::::zt33)` + green + `   "4EEEtttji3P*    ` + "\n" +
			blue + `  :t::::::::tt33.` + yellow + `:Z3z..` + green + `  '' ` + yellow + ` ,..g.    ` + "\n" +
			blue + `  i::::::::zt33F` + yellow + ` AEEEtttt::::ztF     ` + "\n" +
			blue + ` ;:::::::::t33V` + yellow + ` ;EEEttttt::::t3      ` + "\n" +
			blue + ` E::::::::zt33L` + yellow + ` @EEEtttt::::z3F      ` + "\n" +
			blue + `{3=*^'''"*4E3)` + yellow + ` ;EEEtttt:::::tZ'      ` + "\n" +
			blue + `             '` + yellow + ` :EEEEtttt::::z7       ` + "\n" +
			yellow + `                 "VEzjt:;;z>*'       ` + "\n",
	}

	logosData = map[string]string{
		"arch": cyan + `                   -` + "                     %s: %s\n" +
			cyan + `                  o+` + "                     %s: %s\n" +
			cyan + `                 ooo` + "                     %s: %s\n" +
			cyan + `                +oooo:` + "                   %s: %s\n" +
			cyan + `               +oooooo:` + "                  %s: %s\n" +
			cyan + `              --+oooooo+:` + "                %s: %s\n" +
			cyan + `             /:-:++oooo+:` + "                %s: %s\n" +
			cyan + `            /++++/+++++++:` + "               %s: %s\n" +
			cyan + `           /++++++++++++++:` + "              %s: %s\n" +
			cyan + `          /+++o` + blue + `oooooooo` + cyan + `oooo/` + "\n" +
			blue + `         ` + cyan + `./` + blue + `ooosssso++osssssso` + cyan + `+` + "\n" +
			blue + `        .oossssso-\\\\/ossssss+` + "\n" +
			blue + `       -osssssso.      :ssssssso.` + "\n" +
			blue + `      :osssssss/        osssso+++.` + "\n" +
			blue + `     /ossssssss/        +ssssooo/-` + "\n" +
			blue + `    /ossssso+/:-        -:/+osssso+-` + "\n" +
			blue + `   +sso+:-\                 \.-/+oso:` + "\n" +
			blue + `  ++:.                           \-/+/` + "\n" +
			blue + ` .\                                 \/` + "\n",
		"debian": red + `       _,met$$$$$gg.           ` + "\n" +
			red + `    ,g$$$$$$$$$$$$$$$P.       ` + " %s: %s\n" +
			red + `  ,g$$P""       """Y$$.".     ` + " %s: %s\n" +
			red + ` ,$$P'              '$$$.     ` + " %s: %s\n" +
			red + `',$$P       ,ggs.     '$$b:   ` + " %s: %s\n" +
			red + `'d$$'     ,$P"'   ` + white + `.` + red + `    $$$    ` + " %s: %s\n" +
			red + ` $$P      d$'     ` + white + `,` + red + `    $$P    ` + " %s: %s\n" +
			red + ` $$:      $$.   ` + white + `-` + red + `    ,d$$'    ` + " %s: %s\n" +
			red + ` $$\;      Y$b._   _,d$P'     ` + " %s: %s\n" +
			red + ` Y$$.    ` + white + `'.` + red + `'"Y$$$$P"'         ` + " %s: %s\n" +
			red + ` '$$b      ` + white + `"-.__              ` + "\n" +
			red + `  'Y$$                        ` + "\n" +
			red + `   'Y$$.                      ` + "\n" +
			red + `     '$$b.                    ` + "\n" +
			red + `       'Y$$b.                 ` + "\n" +
			red + `          '"Y$b._             ` + "\n" +
			red + `              '""""           ` + "\n",
		"ubuntu": red + `                          ./+o+-       ` + "  %s: %s\n" +
			white + `                  yyyyy- ` + red + `-yyyyyy+     ` + "   %s: %s\n" +
			white + `               ` + white + `://+//////` + red + `-yyyyyyo     ` + "   %s: %s\n" +
			yellow + `           .++ ` + white + `.:/++++++/-` + red + `.+sss/\     ` + "   %s: %s\n" +
			yellow + `         .:++o:  ` + white + ` +/++++++++/:--:/-     ` + " %s: %s\n" +
			yellow + `        o:+o+:++. ` + white + `'..'''.-/oo+++++/    ` + "  %s: %s\n" +
			yellow + `       .:+o:+o/.` + white + `         '+sssoo+/   ` + "    %s: %s\n" +
			white + `  .++/+:` + yellow + `+oo+o:` + white + `             /sssooo.  ` + "    %s: %s\n" +
			white + ` /+++//+:` + yellow + `'oo+o` + white + `               /::--:.  ` + "   %s: %s\n" +
			white + ` \+/+o+++` + yellow + `'o++o` + red + `               ++////.  ` + "\n" +
			white + `  .++.o+` + yellow + ` ++oo+:'` + red + `             /dddhhh.  ` + "\n" +
			yellow + `       .+.o+oo:.` + red + `          \'oddhhhh+   ` + "\n" +
			yellow + `        \+.++o+o\'` + red + `'-''''.:ohdhhhhh+    ` + "\n" +
			yellow + `         ':o+++ ` + red + `'ohhhhhhhhyo++os:     ` + "\n" +
			yellow + `           .o:` + red + `'.syhhhhhhh/` + yellow + `.oo++o'     ` + "\n" +
			red + `               /osyyyyyyo` + yellow + `++ooo+++/    ` + "\n" +
			red + `                   ''''' ` + yellow + `+oo+++o\:    ` + "\n" +
			yellow + `                          'oo++.      ` + "\n",
		"mint": green + `MMMMMMMMMMMMMMMMMMMMMMMMMmds+.       ` + "\n" +
			green + `MMm----::-://////////////oymNMd+'    ` + "\n" +
			green + `MMd      ` + white + `/++                ` + green + `-sNMd:   ` + "\n" +
			green + `MMNso/'  ` + white + `dMM    '.::-. .-::.' ` + green + `.hMN:  ` + "\n" +
			green + `ddddMMh  ` + white + `dMM   :hNMNMNhNMNMNh: ` + green + `'NMm  ` + "\n" +
			green + `    NMm  ` + white + `dMM  .NMN/-+MMM+-/NMN' ` + green + `dMM  ` + "\n" +
			green + `    NMm  ` + white + `dMM  -MMm  'MMM   dMM. ` + green + `dMM  ` + "\n" +
			green + `    NMm  ` + white + `dMM  -MMm  'MMM   dMM. ` + green + `dMM  ` + "\n" +
			green + `    NMm  ` + white + `dMM  .mmd  'mmm   yMM. ` + green + `dMM  ` + "\n" +
			green + `    NMm  ` + white + `dMM'  ..'   ...   ydm. ` + green + `dMM  ` + "\n" +
			green + `    hMM- ` + white + `+MMd/-------...-:sdds  ` + green + `dMM  ` + "\n" +
			green + `    -NMm- ` + white + `:hNMNNNmdddddddddy/'  ` + green + `dMM  ` + "\n" +
			green + `     -dMNs-` + white + `''-::::-------.''    ` + green + `dMM  ` + "\n" +
			green + `      '/dMNmy+/:-------------:/yMMM  ` + "\n" +
			green + `         ./ydNMMMMMMMMMMMMMMMMMMMMM  ` + "\n" +
			green + `            \.MMMMMMMMMMMMMMMMMMM    ` + "\n",
		"centos": yellow +
			`                 .PLTJ.                  ` + "  %s: %s\n" +
			yellow + `                <><><><>                 ` + "  %s: %s\n" +
			green +
			`      KKSSV' 4KKK  ` + yellow + `LJ ` + red + ` KKKL.'VSSKK ` + "        %s: %s\n" +
			green + `      KKV' 4KKKKK ` + yellow + ` LJ` + red + `  KKKKAL 'VKK` + "         %s: %s\n" +
			green + `      V' ' 'VKKKK ` + yellow + ` LJ ` + red + ` KKKKV' ' 'V` + "         %s: %s\n" +
			green + `      .4MA.' 'VKK ` + yellow + ` LJ ` + red + ` KKV' '.4Mb.` + "         %s: %s\n" +
			red + `    . ` + green + `KKKKKA.' VK ` + yellow + ` LJ ` + red + ` V' '.4KKKKK` + blue + `.` + "        %s: %s\n" +
			red + `  .4D ` + green + `KKKKKKKA.'' ` + yellow + ` LJ ` + red + ` ''.4KKKKKKK ` + blue + `FA.` + "     %s: %s\n" +
			red + ` <QDD ++++++++++++  ` + blue + `  ++++++++++++ GFD>` + "    %s: %s\n" +
			red + `  'VD ` + blue + `KKKKKKKK'.. ` + green + ` LJ ` + yellow + ` ..'KKKKKKKK ` + blue + `FV` + "\n" +
			red + `    ' ` + blue + `VKKKKK'. .4 ` + green + ` LJ  ` + yellow + `K. .'KKKKKV` + blue + `'` + "\n" +
			blue + `      'VK'. .4KK   ` + green + `LJ ` + yellow + ` KKA. .'KV'` + "\n" +
			blue + `      A. . .4KKKK  ` + green + `LJ  ` + yellow + `KKKKA. . .4` + "\n" +
			blue + `      KKA. 'KKKKK  ` + green + `LJ  ` + yellow + `KKKKK' .4KK` + "\n" +
			blue + `      KKSSA. VKKK  ` + green + `LJ  ` + yellow + `KKKV .4SSKK` + "\n" +
			green + `                <><><><>` + "\n" +
			green + `                 'MKKM'` + "\n",
		"osx": green + `                  -/+:.          ` + "\n" +
			green + `                 :++++.         ` + "\n" +
			green + `                /+++/.          ` + "\n" +
			green + `        .:-::- .+/:-''.::-      ` + "\n" +
			green + `     .:/++++++/::::/++++++/:'   ` + "\n" +
			yellow + `  .:///////////////////////:'  ` + "\n" +
			yellow + `  ////////////////////////'    ` + "\n" +
			orange + ` -+++++++++++++++++++++++'     ` + "\n" +
			orange + ` /++++++++++++++++++++++/      ` + "\n" +
			red + ` /sssssssssssssssssssssss.     ` + "\n" +
			red + ` :ssssssssssssssssssssssss-    ` + "\n" +
			purple + `  osssssssssssssssssssssssso/' ` + "\n" +
			purple + `  'syyyyyyyyyyyyyyyyyyyyyyyy+' ` + "\n" +
			blue + `   'ossssssssssssssssssssss/   ` + "\n" +
			blue + `     :ooooooooooooooooooo+.    ` + "\n" +
			blue + `      ':+oo+/:-..-:/+o+/-      ` + "\n",
		"windows": red + `        ,.=:!!t3Z3z.,                 ` + "\n" +
			red + `       :tt:::tt333EE3                ` + "\n" +
			red + `       Et:::ztt33EEEL` + green + ` @Ee.,      .., ` + "\n" +
			red + `      ;tt:::tt333EE7` + green + ` ;EEEEEEttttt33# ` + "\n" +
			red + `     :Et:::zt333EEQ.` + green + ` $EEEEEttttt33QL ` + "\n" +
			red + `     it::::tt333EEF` + green + ` @EEEEEEttttt33F  ` + "\n" +
			red + `    ;3=*^'''"*4EEV` + green + ` :EEEEEEttttt33@.  ` + "\n" +
			blue + `    ,.=::::!t=., ` + red + `'` + green + ` @EEEEEEtttz33QF   ` + "\n" +
			blue + `   ;::::::::zt33)` + green + `   "4EEEtttji3P*    ` + "\n" +
			blue + `  :t::::::::tt33.` + yellow + `:Z3z..` + green + `  '' ` + yellow + ` ,..g.    ` + "\n" +
			blue + `  i::::::::zt33F` + yellow + ` AEEEtttt::::ztF     ` + "\n" +
			blue + ` ;:::::::::t33V` + yellow + ` ;EEEttttt::::t3      ` + "\n" +
			blue + ` E::::::::zt33L` + yellow + ` @EEEtttt::::z3F      ` + "\n" +
			blue + `{3=*^'''"*4E3)` + yellow + ` ;EEEtttt:::::tZ'      ` + "\n" +
			blue + `             '` + yellow + ` :EEEEtttt::::z7       ` + "\n" +
			yellow + `                 "VEzjt:;;z>*'       ` + "\n",
	}

	colorsData = map[string]string{
		"arch":   "blue",
		"ubuntu": "red",
		"debian": "red",
		"centos": "yellow",
		"osx":    "yellow",
	}
)

// GetOperatingSystems returns available logos
func GetOperatingSystems() []string {
	keys := make([]string, 0, len(logos))
	for key := range logos {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

// GetLogo return logo which match distribution's name
func GetLogo(name string) string {
	return logos[name]
}

// GetOperatingSystemTheme return logo which match distribution's name
// with formatting output and main color
func GetOperatingSystemTheme(name string) (string, string, error) {
	logo, okLogo := logosData[name]
	if !okLogo {
		return "", "", fmt.Errorf("Unsupported operating system: %s", name)
	}

	color, okColor := colorsData[name]
	if !okColor {
		color = "white"
	}
	return logo, color, nil
}
