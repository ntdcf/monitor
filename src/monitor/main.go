package main

import (
	"log"
	"monitor/setting"
	"monitor/work"
)

func main() {
	configuration := new(setting.Configuration)
	configuration.Init()
	log.Println(configuration.MonitorConfigs)
	monitors := configuration.MonitorConfigs.Configures
	for _, monitorConfig := range monitors {
		switch monitorConfig.Monitor {
		case "cpu":
			cpu := &work.Cpu{}
			cpu.GetMonitorData()
			log.Println(cpu)
			break
		case "memory":
			log.Println(monitorConfig)
			break
		default:
			log.Println("不受支持的监控类型")

		}
	}
}
