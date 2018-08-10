package work

import (
	"time"
)

type Monitor interface {
	SetMonitorData()
	// ReportMonitorData(data *MonitorData)
}

type MonitorData struct {
	Data        float64
	MonitorTime time.Time
}
