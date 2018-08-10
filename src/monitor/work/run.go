package work

import (
	"log"
	"time"
)

func createMonitor(monitorType string) Monitor {
	switch monitorType {
	case "cpu":
		return &Cpu{}
	case "disk":
		return &Disk{}
	case "free":
		log.Println(111)
		return &Disk{}
	case "uptime":
		return nil
	default:
		log.Println("不受支持的监控类型")
		return nil
	}
}

func Run(monitotType string, circle int64, stop chan bool) {
	monitor := createMonitor(monitotType)

	if monitor != nil {
		for {
			monitor.SetMonitorData()
			log.Println(monitor)
			select {
			case flag := <-stop:
				if flag == true {
					return
				}
			default:
			}
			time.Sleep(time.Duration(circle) * time.Second)
		}
	}
}
