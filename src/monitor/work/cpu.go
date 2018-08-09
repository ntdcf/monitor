package work

import (
	"time"
)

type Cpu struct {
	MonitorData
}

func (c *Cpu) GetMonitorData() {
	c.Data = 100
	c.MonitorTime = time.Now()
}
