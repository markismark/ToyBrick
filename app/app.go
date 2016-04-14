package app

import (
	"fmt"
	"log"
	"net/http"
)

func Run() {
	http.HandleFunc("/", resquestHandler)
	err := http.ListenAndServe(":8123", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
	fmt.Println("begin run")

}

func resquestHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/favicon.ico" {
		return
	}
	r.ParseForm()
	data := r.Form["data"]
	if len(data) == 0 {
		return
	}
	textRequestJson := data[0]
	tr := createTextRequest(textRequestJson, r)
	hr, _ := BuildHttpRequest(tr)
	response, err := SendRequest(hr)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	fmt.Fprintf(w, response.Body)
}

func loadConfig() {

}

func parseParams() {

}
