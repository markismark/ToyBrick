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
	contentByte, _ := json.Marshal(responses)
	content := string(contentByte)

	callback := ""
	callbackArr := r.Form["callback"]
	if len(callbackArr) != 0 {
		callback = callbackArr[0]
		content = fmt.Sprintf("if (window.%s)%s(%s)", callback, callback, content)
		w.Header().Set("Content-Type", "application/javascript")
	}

	fmt.Fprintf(w, "%s", content)

}

func loadConfig() {

}

func parseParams() {

}
