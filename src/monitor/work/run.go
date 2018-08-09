package work

func Run(monitor Monitor) {
	data := monitor.GetMonitorData()
	monitor.ReportMonitorData(data)
}
