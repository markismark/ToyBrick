package app

import (
	"net/http"
)

type TextResponse struct {
	Body       string
	HttpStatus int
	SetCookie  map[string]string
}

func BuildTextResponse(response http.Response) {

}
