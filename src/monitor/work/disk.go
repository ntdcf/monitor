package work

import (
	"syscall"
	"time"
)

type Disk struct {
	MonitorData
}

func (d *Disk) SetMonitorData() {
	fs := syscall.Statfs_t{}
	err := syscall.Statfs("/", &fs)
	if err != nil {
		return
	}
	free := float64(fs.Bfree)
	all := float64(fs.Blocks)
	used := all - free
	d.Data = used / all
	d.MonitorTime = time.Now()
}
