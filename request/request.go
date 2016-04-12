package request

import (
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

func BuildHttpRequest(tr *TextRequest) *http.Request {

}
