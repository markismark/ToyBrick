package conf

import (
	"log"
	"mi_com_tool_dataset/util"
	"strconv"
	"strings"

	"github.com/Maxgis/tree"
	"github.com/go-ini/ini"
)

var (
	Tags                    = make(map[string]*Group)
	DEFAULE_TIMEOUT  uint64 = 10000
	DEFAULT_PROTOCOL        = "http"
	DEFAULT_MAXRETRY uint   = 1
)

type Group struct {
	Timeout  uint64
	Protocol string
	Balance  string
	MaxRetry uint
	Host     string
	Machines *[]*Machine
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
		//log.Printf("%#v\n", Tags[section.Name()])
		if Tags[section.Name()].Machines == nil {
			continue
		}
		for _, machine := range *Tags[section.Name()].Machines {
			log.Printf("%#v\n", machine)
		}
	}

	tree.Print(Tags)
}

func initGroup(section *ini.Section) *Group {
	timeout := getTimeout(section)
	protocol := getProtocol(section)
	maxRetry := getMaxRetry(section)
	host := getHost(section)
	machines := getMachines(section)
	return &Group{
		Timeout:  timeout,
		Protocol: protocol,
		MaxRetry: maxRetry,
		Host:     host,
		Machines: machines,
	}

}

func getTimeout(section *ini.Section) uint64 {
	var err error
	var timeoutKey *ini.Key
	timeoutKey, err = section.GetKey("timeout")
	if err != nil {
		return DEFAULE_TIMEOUT
	}
	return timeoutKey.MustUint64(DEFAULE_TIMEOUT)
}

func getProtocol(section *ini.Section) string {
	var err error
	var protocolKey *ini.Key
	protocolKey, err = section.GetKey("protocol")
	if err != nil {
		return DEFAULT_PROTOCOL
	}
	protocol := protocolKey.String()
	if protocol == "" {
		return DEFAULT_PROTOCOL
	}
	return protocol
}

func getMaxRetry(section *ini.Section) uint {
	var err error
	var maxRetryKey *ini.Key
	maxRetryKey, err = section.GetKey("max_retry")
	if err != nil {
		return DEFAULT_MAXRETRY
	}
	maxRetry := maxRetryKey.MustUint(DEFAULT_MAXRETRY)
	return maxRetry
}

func getHost(section *ini.Section) string {
	var err error
	var hostKey *ini.Key
	hostKey, err = section.GetKey("host")
	if err != nil {
		return ""
	}
	return hostKey.MustString("")
}

func getMachines(section *ini.Section) *[]*Machine {
	var err error
	var machinesKey *ini.Key
	machinesKey, err = section.GetKey("machines")
	if err != nil {
		return nil
	}
	machinesStr := machinesKey.MustString("")
	machinesArr := strings.Split(machinesStr, ",")
	length := len(machinesArr)
	if length == 0 {
		return nil
	}
	machines := make([]*Machine, length)
	for i, machineStr := range machinesArr {
		machineInfo := strings.Split(machineStr, ":")
		var port int
		if length > 1 {
			port, err = strconv.Atoi(machineInfo[1])
		}
		machines[i] = &Machine{Host: machineInfo[0], Port: port}
	}
	return &machines
}
