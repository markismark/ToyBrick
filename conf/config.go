package conf

import (
	"log"
	"strings"

	"github.com/Maxgis/ToyBrick/util"
	"github.com/go-ini/ini"
)

type Config struct {
	Port                  string
	IsOpenReferrer        bool
	ReferrerWhiteList     []string
	IsOpenDomainWhitelist bool
	DomainWhitelist       []string
	IsOpenAdmin           bool
	AdminPort             string
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
	Globals.DomainWhitelist = GetOpenDomainWhitelist()
	Globals.IsOpenDomainWhitelist = IsOpenDomainWhitelist()
	Globals.AdminPort = GetAdminPort()
	Globals.IsOpenAdmin = IsOpenAdmin()
}

func GetPort() (string, error) {
	portKey, err := cfg.Section("basic").GetKey("port")
	if err != nil {
		log.Fatal(err)
	}
	return portKey.String(), nil
}

func GetAdminPort() string {
	portKey, err := cfg.Section("admin").GetKey("port")
	if err != nil {
		return ""
	}
	return portKey.String()
}

func IsOpenAdmin() bool {
	isOpenKey, err := cfg.Section("admin").GetKey("open")
	if err != nil {
		return false
	}
	isOpen, perr := isOpenKey.Bool()
	if perr != nil {
		return false
	}
	return isOpen
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

func IsOpenDomainWhitelist() bool {
	isOpenKey, err := cfg.Section("security").GetKey("openDomain")
	if err != nil {
		return false
	}
	isOpen, perr := isOpenKey.Bool()
	if perr != nil {
		return false
	}
	return isOpen
}

func GetOpenDomainWhitelist() []string {
	domainWhiteList, err := cfg.Section("security").GetKey("domainWhiteList")
	if err != nil {
		return []string{}
	}

	return strings.Split(domainWhiteList.String(), ",")
}
