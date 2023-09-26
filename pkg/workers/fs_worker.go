package workers

import (
	"fmt"

	human "github.com/dustin/go-humanize"
	"github.com/shirou/gopsutil/disk"
)

func DriveInfo() {
	formatter := "%-14s %7s %7s %7s %4s %s\n"
	fmt.Printf(formatter, "Filesystem", "Size", "Used", "Avail", "Use%", "Mounted on")

	parts, _ := disk.Partitions(true)
	for _, p := range parts {
		device := p.Mountpoint
		s, _ := disk.Usage(device)

		if s.Total == 0 {
			continue
		}

		percent := fmt.Sprintf("%2.f%%", s.UsedPercent)

		fmt.Printf(formatter,
			s.Fstype,
			human.Bytes(s.Total),
			human.Bytes(s.Used),
			human.Bytes(s.Free),
			percent,
			p.Mountpoint,
		)
	}
}
