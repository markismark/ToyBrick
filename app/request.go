package app

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"time"
)

type TextRequest struct {
	Header  map[string]string
	Method  string
	Cookie  map[string]string
	Data    map[string]map[string]string
	Url     string
	Status  string
	Retry   int
	Timeout time.Duration
}

func BuildHttpRequest(tr *TextRequest) (*http.Request, error) {
	req, err := http.NewRequest(tr.Method, tr.Url, bytes.NewBuffer([]byte{}))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "text/plain")

	return req, nil
}

func SendRequest(hr *http.Request) (*TextResponse, error) {
	client := &http.Client{}
	resp, _ := client.Do(hr)
	body, _ := ioutil.ReadAll(resp.Body)
	tr := &TextResponse{}
	tr.Body = string(body)
	tr.HttpStatus = resp.StatusCode
	return tr, nil
}
