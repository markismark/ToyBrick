package conf

import (
	"log"
	"mi_com_tool_dataset/util"

	"github.com/go-ini/ini"
)

func init() {
	file := util.GetCurrentDirectory() + "/conf/machine.ini"
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
