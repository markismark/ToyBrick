package conf

import (
	"log"
	"strings"

	"github.com/Maxgis/ToyBrick/util"
	"github.com/go-ini/ini"
)

type Config struct {
	Port              string
	IsOpenReferrer    bool
	ReferrerWhiteList []string
}

var (
	cfg     *ini.File
	Globals = &Config{}
)

func init() {
	file := util.GetCurrentDirectory() + "/conf/conf.ini"
	var err error
	cfg, err = ini.LooseLoad("filename", file)
	if err != nil {
		log.Fatal(err)
	}
	port, err := GetPort()
	if err != nil {
		log.Fatal(err)
	}
	Globals.Port = port
	Globals.IsOpenReferrer = IsOpenReferrer()
	Globals.ReferrerWhiteList = GetReferrerWhiteList()
}

func GetPort() (string, error) {
	portKey, err := cfg.Section("basic").GetKey("port")
	if err != nil {
		log.Fatal(err)
	}
	return portKey.String(), nil
}

func IsOpenReferrer() bool {
	isOpenKey, err := cfg.Section("security").GetKey("openReferrer")
	if err != nil {
		return false
	}
	isOpen, perr := isOpenKey.Bool()
	if perr != nil {
		return false
	}
	return isOpen
}

func GetReferrerWhiteList() []string {
	referrerWhiteList, err := cfg.Section("security").GetKey("referrerWhiteList")
	if err != nil {
		return []string{}
	}

	return strings.Split(referrerWhiteList.String(), ",")
}
