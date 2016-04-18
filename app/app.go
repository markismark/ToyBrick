package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//Run function
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
	var trList []TextRequest
	textRequestJSON := data[0]
	err := json.Unmarshal([]byte(textRequestJSON), &trList)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	for _, tr := range trList {
		log.Printf("%#v", tr)
		InitTextRequest(&tr, r)
		hr, _ := BuildHttpRequest(&tr)
		response, err := SendRequest(hr)
		if err != nil {
			fmt.Fprintln(w, err)
			return
		}
		fmt.Fprintf(w, "%s-------------------------\n", response.Body)
	}
	// tr, err := createTextRequest(textRequestJSON, r)
	// hr, _ := BuildHttpRequest(tr)
	// response, err := SendRequest(hr)
	// if err != nil {
	// 	fmt.Fprintln(w, err)
	// 	return
	// }
	//fmt.Fprintf(w, response.Body)
}

func loadConfig() {

}

func parseParams() {

}
