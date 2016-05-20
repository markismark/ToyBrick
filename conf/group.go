package conf

import (
	"log"
	"mi_com_tool_dataset/util"

	"github.com/go-ini/ini"
)

var (
	Tags    = make(map[string]*Group)
	TIMEOUT = 10000
)

type Group struct {
	Timeout  int
	Protocol string
	Balance  string
	MaxRetry int
	Host     string
	Machines *[]Machine
}

type Machine struct {
	Host string
	Port int
}

func init() {
	file := util.GetCurrentDirectory() + "/conf/group.ini"
	var err error
	cfg, err = ini.LooseLoad("filename", file)
	if err != nil {
		log.Fatal(err)
	}
	sections := cfg.Sections()
	for _, section := range sections {
		Tags[section.Name()] = initGroup(section)
		log.Printf("%#v\n", Tags[section.Name()])
	}

	//log.Printf("%#v\n", Tags)
}

func initGroup(section *ini.Section) *Group {
	timeout := getTimeout(section)
	protocol := getProtocol(section)
	return &Group{Timeout: timeout, Protocol: protocol}

}

func getTimeout(section *ini.Section) int {
	var err error
	var timeout int
	var timeoutKey *ini.Key
	timeoutKey, err = section.GetKey("timeout")
	if err != nil {
		timeout = TIMEOUT
	} else {
		timeout, err = timeoutKey.Int()
		if err != nil {
			timeout = TIMEOUT
		}
	}
	return timeout
}

func getProtocol(section *ini.Section) string {
	var err error
	var protocolKey *ini.Key
	protocolKey, err = section.GetKey("protocol")
	if err != nil {
		return ""
	}
	return protocolKey.String()
}
