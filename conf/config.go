package conf

import (
	"log"

	"github.com/Maxgis/ToyBrick/util"
	"github.com/go-ini/ini"
)

var cfg *ini.File

func init() {
	file := util.GetCurrentDirectory() + "/conf/conf.ini"
	var err error
	cfg, err = ini.LooseLoad("filename", file)
	if err != nil {
		log.Fatal(err)
	}
}

func GetPort() (string, error) {
	port, err := cfg.Section("basic").GetKey("port")
	if err != nil {
		log.Fatal(err)
	}
	return port.String(), nil
}
