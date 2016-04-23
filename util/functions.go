package util

import (
	"log"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}

func GetUrlDomain(urlStr string) string {
	urlInfo, err := url.Parse(urlStr)
	if err != nil {
		return ""
	}
	return urlInfo.Host
}

func HostIsInList(host string, list []string) bool {
	for _, h := range list {
		if HostIsMatch(host, h) {
			return true
		}
	}
	return false
}

func HostIsMatch(host string, pattern string) bool {
	if strings.HasPrefix(pattern,"*") {
		return strings.HasSuffix(host, pattern[1:])
	}
	return host == pattern
}
