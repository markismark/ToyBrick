package app

import (
	"net/http"
)

type TextResponse struct {
	Body       string `json:"body"`
	HttpStatus int    `json:"status"`
	//SetCookie  map[string]string `json:"cookie"`
}

func BuildTextResponse(response http.Response) {

}
