package app

import (
	"net/http"
)

func Run() {
	http.Handle("/", resquestHandler)
	http.ListenAndServe("8123", nil)
}

func resquestHandler(w http.ResponseWriter, r *http.Request) {

}

func loadConfig() {

}
