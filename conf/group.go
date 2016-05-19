package conf

import (
	"log"
	"mi_com_tool_dataset/util"

	"github.com/go-ini/ini"
)

type Group struct {
	Name     string
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
	tags := cfg.Sections()
	for _, tag := range tags {
		log.Printf("%#v", tag.Name())
	}
}
