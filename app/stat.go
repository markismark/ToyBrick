package app

import (
	"log"
	"net/url"
	"sync"
	"time"
)

var StatisticsMap = &HostMap{hostmap: make(map[string]*URLMap)}

type Statistics struct {
	Path       string
	RequestNum int64
	MinTime    time.Duration
	MaxTime    time.Duration
	TotalTime  time.Duration
}

type URLMap struct {
	Host       string
	RequsetNum int64
	lock       sync.RWMutex
	pathmap    map[string]*Statistics
}

func (this *URLMap) addPath(path string) *Statistics {
	stat := &Statistics{Path: path}
	this.pathmap[path] = stat
	return stat
}

type HostMap struct {
	hostmap map[string]*URLMap
}

func (this *HostMap) addHost(host string) *URLMap {
	urlmap := &URLMap{
		pathmap:    make(map[string]*Statistics),
		RequsetNum: 0,
		Host:       host,
	}
	this.hostmap[host] = urlmap
	return urlmap
}

func (this *HostMap) AddStatistic(URL *url.URL) {
	host := URL.Host
	path := URL.Path
	if path == "" {
		path = "/"
	}
	urlmap, ok := StatisticsMap.hostmap[host]
	if !ok {
		urlmap = this.addHost(host)
	}
	urlmap.lock.Lock()
	defer urlmap.lock.Unlock()
	urlmap.RequsetNum++

	pathStat, pok := urlmap.pathmap[path]
	if !pok {
		pathStat = urlmap.addPath(path)
	}
	pathStat.RequestNum++
	log.Printf("%s, %#v\n", host, urlmap)
	log.Printf("%s, %#v\n", path, pathStat)
}

func AddStat(URL *url.URL) {
	StatisticsMap.AddStatistic(URL)
}
