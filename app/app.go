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
	chs := []chan int{}
	responses := make([]*TextResponse, len(trList))
	for i, tr := range trList {
		log.Printf("%#v", tr)
		InitTextRequest(&tr, r)
		hr, _ := BuildHttpRequest(&tr)
		ch := make(chan int)
		chs = append(chs, ch)
		go func(i int) {
			response, err := SendRequest(hr)
			log.Printf("get response %d\n", i)
			responses[i] = response
			if err != nil {
				fmt.Fprintln(w, err)
			}
			ch <- 0
		}(i)
	}
	for _, ch := range chs {
		<-ch
	}
	log.Println("get all response")
	for _, response := range responses {
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
