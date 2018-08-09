package setting

import (
	"encoding/xml"
	"io/ioutil"
	"log"
)

type Configuration struct {
	MonitorConfigs *Configs
}

func (c *Configuration) Init() {
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}
	c.getMonitorConfig(file)
}

func (c *Configuration) getMonitorConfig(file []byte) {
	configs := &Configs{}
	err := xml.Unmarshal(file, configs)
	if err != nil {
		log.Fatal(err)
	}
	c.MonitorConfigs = configs
}
