package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Maxgis/ToyBrick/conf"
	"github.com/Maxgis/ToyBrick/util"
)

//Run function
func Run() {
	http.HandleFunc("/proxy", resquestHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./libs"))))
	port, err := conf.GetPort()
	if err != nil {
		log.Fatal(err)
	}
	herr := http.ListenAndServe(":"+port, nil)
	if herr != nil {
		log.Fatal("ListenAndServe:", herr)
	}
	fmt.Println("begin run")

}

func resquestHandler(w http.ResponseWriter, r *http.Request) {

	if conf.IsOpenReferrer() {
		referer := r.Header.Get("referer")
		refererHost := util.GetUrlDomain(referer)
		refererList := conf.GetReferrerWhiteList()
		if !util.HostIsInList(refererHost, refererList) {
			fmt.Fprintln(w, "over")
			return
		}
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
	} else {
		w.Header().Set("Content-Type", "application/json")
	}

	fmt.Fprintf(w, "%s", content)

}
