package work

import (
	"time"
)

type Monitor interface {
	GetMonitorData() *MonitorData
	ReportMonitorData(data *MonitorData)
}

type MonitorData struct {
	Data        int64
	MonitorTime time.Time
}
