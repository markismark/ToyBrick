package app

import (
	"compress/gzip"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/Maxgis/ToyBrick/conf"
)

type TextRequest struct {
	Header      map[string]string `json:"header"`
	Method      string            `json:"method"`
	Cookies     map[string]string `json:"cookies"`
	Data        string            `json:"data"`
	URL         string            `json:"url"`
	Status      string            `json:-`
	Retry       int               `json:"retry"`
	Timeout     int               `json:"timeout"`
	UserRequset *http.Request     `json:-`
	Tag         string            `json:tag`
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
		"Accept-Encoding",
		//"Connection",
	}
)

func InitTextRequest(tr *TextRequest, hr *http.Request) {
	tr.Method = strings.ToUpper(tr.Method)
	tr.UserRequset = hr
}

func BuildHttpRequest(tr *TextRequest) (*http.Request, error) {

	if tr.Tag != "" {
		tr.URL = conf.BuildURI(tr.Tag, tr.URL)
	}

	req, err := http.NewRequest(tr.Method, tr.URL, strings.NewReader(tr.Data))

	if err != nil {
		return nil, err
	}

	for _, key := range UserHeader {
		if value := tr.UserRequset.Header.Get(key); value != "" {
			req.Header.Set(key, value)
		}
	}
	if req.Method == "POST" {
		//If not set content-type, some servers can't parse the data.
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	//tr.UserRequset.Header.Set("Connection", "keep-alive")
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
	var body string
	tr := &TextResponse{}
	buf := make([]byte, 1024)
	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, _ := gzip.NewReader(resp.Body)
		defer reader.Close()
		for {
			n, err := reader.Read(buf)
			if err != nil && err != io.EOF {
				return nil, err
			}
			if n == 0 {
				break
			}
			body += string(buf[:n])
		}
	default:
		bodyByte, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		body = string(bodyByte)
	}

	tr.Body = body
	tr.HttpStatus = resp.StatusCode
	return tr, nil
}
