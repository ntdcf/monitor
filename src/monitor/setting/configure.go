package setting

import (
	"encoding/xml"
)

const configPath = "conf/config.xml"

type Configs struct {
	XMLName    xml.Name    `xml:"configures"`
	Configures []Configure `xml:"configure"`
}

type Configure struct {
	XMLName xml.Name `xml:"configure"`
	Monitor string   `xml:"monitor"`
	Circle  int64    `xml:"circle"`
}
