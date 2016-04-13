package app

import (
	"fmt"
	"net/http"
)

func Run() {
	//	http.HandleFunc("/", resquestHandler)
	//	err := http.ListenAndServe(":8123", nil)
	//	if err != nil {
	//		log.Fatal("ListenAndServe:", err)
	//	}
	//	fmt.Println("begin run")
	tr := &TextRequest{Method: "post", Url: "http://buy.mi.com/in/accessories"}
	hr, _ := BuildHttpRequest(tr)
	response, _ := SendRequest(hr)
	fmt.Println(response)
}

func resquestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello ! welcome")
}

func loadConfig() {

}

func parseParams() {

}
