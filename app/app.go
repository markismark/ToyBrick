package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/Maxgis/ToyBrick/conf"
	"github.com/Maxgis/ToyBrick/util"
)

//Run function
func Run() {
	http.HandleFunc("/proxy", resquestHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./libs"))))
	log.Println("start proxy ,port:", conf.Globals.Port)
	err := http.ListenAndServe(":"+strconv.Itoa(conf.Globals.Port), nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
	fmt.Println("begin run")

}

func resquestHandler(w http.ResponseWriter, r *http.Request) {

	if !checkReferer(r) {
		fmt.Fprintln(w, "over")
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
		if !checkDomain(tr.URL) {
			responses[i] = &TextResponse{HttpStatus: ERROR_NOT_IN_WHITELIST}
			continue
		}
		InitTextRequest(&tr, r)
		hr, _ := BuildHttpRequest(&tr)
		ch := make(chan int)
		chs = append(chs, ch)
		go func(i int) {
			response, err := SendRequest(hr)
			log.Printf("get response %d\n", i)
			responses[i] = response
			AddStat(hr.URL)
			if err != nil {
				fmt.Fprintln(w, err)
			}
			ch <- 0
		}(i)
	}
	for _, ch := range chs {
		<-ch
		close(ch)
	}
	contentByte, _ := json.Marshal(responses)
	content := string(contentByte)

	callback := ""
	callbackArr := r.Form["callback"]
	if len(callbackArr) != 0 {
		callback = callbackArr[0]
		reg := regexp.MustCompile(`^[0-9a-zA-Z_]*$`)
		if !reg.MatchString(callback) {
			fmt.Fprintln(w, "callback params error")
			return
		}
		content = fmt.Sprintf("if (window.%s)%s(%s)", callback, callback, content)
		w.Header().Set("Content-Type", "application/javascript")
	} else {
		w.Header().Set("Content-Type", "application/json")
	}

	fmt.Fprintf(w, "%s", content)

}

func checkReferer(r *http.Request) bool {
	if conf.Globals.IsOpenReferrer {
		referer := r.Header.Get("referer")
		refererHost := util.GetUrlDomain(referer)
		if !util.HostIsInList(refererHost, conf.Globals.ReferrerWhiteList) {
			return false
		}
	}
	return true
}

func checkDomain(url string) bool {
	if conf.Globals.IsOpenDomainWhitelist {
		requertHost := util.GetUrlDomain(url)
		if !util.HostIsInList(requertHost, conf.Globals.DomainWhitelist) {
			return false
		}
	}
	return true
}
