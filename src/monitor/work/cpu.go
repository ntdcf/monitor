package work

import (
	"time"
)

type Cpu struct {
	MonitorData
}

func (c *Cpu) SetMonitorData() {
	c.Data = 100
	c.MonitorTime = time.Now()
}
