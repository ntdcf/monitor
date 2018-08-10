package main

import (
	"monitor/setting"
	"monitor/work"
	"time"
)

func main() {
	configuration := new(setting.Configuration)
	configuration.Init()

	monitors := configuration.MonitorConfigs.Configures
	count := len(monitors)
	stop := make([]chan bool, count)
	for i := 0; i < count; i++ {
		stop[i] = make(chan bool, 1)
	}

	for {
		index := 0
		for _, monitorConfig := range monitors {
			go work.Run(monitorConfig.Monitor, monitorConfig.Circle, stop[index])
			index++
		}
		time.Sleep(5 * time.Second)
		stop[0] <- true
		time.Sleep(100 * time.Second)
	}
}
