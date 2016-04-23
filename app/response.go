package app

import (
	"net/http"
)

var (
	ERROR_NOT_IN_WHITELIST int = 1001
)

type TextResponse struct {
	Body       string `json:"body"`
	HttpStatus int    `json:"status"`
	//SetCookie  map[string]string `json:"cookie"`
}

func BuildTextResponse(response http.Response) {

}
