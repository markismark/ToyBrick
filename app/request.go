package app

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type TextRequest struct {
	Header      map[string]string            `json:header`
	Method      string                       `json:method`
	Cookies     map[string]string            `json:cookies`
	Data        map[string]map[string]string `json:data`
	Url         string                       `json:url`
	Status      string                       `json:-`
	Retry       int                          `json:retry`
	Timeout     int                          `json:timeout`
	UserRequset *http.Request                `json:-`
}

var (
	UserHeader = []string{
		"Content-Type",
		"Accept",
		"Accept-Language",
		"User-Agent",
		"Referer",
		"Cache-Control",
		"Cookie",
		"If-Modified-Since",
		"Etag",
	}
)

func createTextRequest(jsonStr string, hr *http.Request) *TextRequest {
	tr := &TextRequest{}
	json.Unmarshal([]byte(jsonStr), tr)
	tr.Method = strings.ToUpper(tr.Method)
	tr.UserRequset = hr
	return tr
}

func BuildHttpRequest(tr *TextRequest) (*http.Request, error) {
	req, err := http.NewRequest(tr.Method, tr.Url, nil)
	if err != nil {
		return nil, err
	}

	for _, key := range UserHeader {
		if value := tr.UserRequset.Header.Get(key); value != "" {
			req.Header.Set(key, value)
		}
	}
	return req, nil
}

func SendRequest(hr *http.Request) (*TextResponse, error) {
	client := &http.Client{}
	resp, perr := client.Do(hr)

	if perr != nil {
		return nil, perr
	}
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
			return
		} else {
			resp.Body.Close()
		}
	}()
	log.Println(resp)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	tr := &TextResponse{}
	tr.Body = string(body)
	tr.HttpStatus = resp.StatusCode
	return tr, nil
}
