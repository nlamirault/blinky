// Copyright (C) 2015 Nicolas Lamirault <nicolas.lamirault@gmail.com>

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
	"bufio"
	"os"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"golang.org/x/sys/unix"

	"github.com/nlamirault/blinky/utils"
)

// Date holds the date facts.
type Date struct {
	Unix int64
	UTC  string
}

// FileSystems holds the Filesystem facts.
type FileSystems map[string]FileSystem

type OSSystem struct {
	Domainname   string
	Architecture string
	Hostname     string
}

// type SystemFacts struct {
// 	BootID      string
// 	Date        Date
// 	OSSystem    OSSystem
// 	Network     Network
// 	Kernel      Kernel
// 	MachineID   string
// 	Memory      Memory
// 	OSRelease   OSRelease
// 	Uptime      int64
// 	FileSystems FileSystems
// }

// Network holds the network facts.
type Network struct {
	Interfaces Interfaces
}

// Interfaces holds the interface facts.
type Interfaces map[string]Interface

// Interface holds facts for a single interface.
type Interface struct {
	Name         string
	Index        int
	HardwareAddr string
	IpAddresses  []string
}

// OSRelease holds the OS release facts.
type OSRelease struct {
	Name       string
	ID         string
	PrettyName string
	Version    string
	VersionID  string
}

// Kernel holds the kernel facts.
type Kernel struct {
	Name    string
	Release string
	Version string
}

// Memory holds the memory facts.
type Memory struct {
	Total    uint64
	Free     uint64
	Shared   uint64
	Buffered uint64
}

// FileSystem holds facts for a filesystem (man fstab).
type FileSystem struct {
	Device     string
	MountPoint string
	Type       string
	Options    []string
	DumpFreq   uint64
	PassNo     uint64
}

func GetDate() *Date {
	now := time.Now()
	return &Date{
		Unix: now.Unix(),
		UTC:  now.UTC().String(),
	}
}

func GetOSRelease() (*OSRelease, error) {
	osReleaseFile, err := os.Open("/etc/os-release")
	if err != nil {
		log.Errorf("Can't open OS release file : %s", err.Error())
		return nil, err
	}
	defer osReleaseFile.Close()
	osrelease := new(OSRelease)

	scanner := bufio.NewScanner(osReleaseFile)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			columns := strings.Split(line, "=")
			key := columns[0]
			value := strings.Trim(strings.TrimSpace(columns[1]), `"`)
			switch key {
			case "NAME":
				osrelease.Name = value
			case "ID":
				osrelease.ID = value
			case "PRETTY_NAME":
				osrelease.PrettyName = value
			case "VERSION":
				osrelease.Version = value
			case "VERSION_ID":
				osrelease.VersionID = value
			}
		}
	}
	return osrelease, nil
}

func GetKernelInformations() (*OSSystem, *Kernel, error) {
	var buf unix.Utsname
	err := unix.Uname(&buf)
	if err != nil {
		return nil, nil, err
	}
	ossystem := new(OSSystem)
	ossystem.Domainname = utils.CharsToString(buf.Domainname)
	ossystem.Architecture = utils.CharsToString(buf.Machine)
	ossystem.Hostname = utils.CharsToString(buf.Nodename)
	kernel := new(Kernel)
	kernel.Name = utils.CharsToString(buf.Sysname)
	kernel.Release = utils.CharsToString(buf.Release)
	kernel.Version = utils.CharsToString(buf.Version)
	return ossystem, kernel, nil
}

func GetLoadAndMemory() (*Memory, int64, error) {
	var info unix.Sysinfo_t
	if err := unix.Sysinfo(&info); err != nil {
		return nil, 0, err
	}

	memory := new(Memory)
	memory.Total = info.Totalram
	memory.Free = info.Freeram
	memory.Shared = info.Sharedram
	memory.Buffered = info.Bufferram
	return memory, info.Uptime, nil
}
