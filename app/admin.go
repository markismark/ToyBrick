package app

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/Maxgis/ToyBrick/conf"
)

var mux map[string]func(http.ResponseWriter, *http.Request)

func init() {
	if conf.Globals.IsOpenAdmin {
		if conf.Globals.AdminPort == 0 {
			return
		}
		server := http.Server{
			Addr:    ":" + strconv.Itoa(conf.Globals.AdminPort),
			Handler: &AdminHandle{},
		}
		mux = make(map[string]func(http.ResponseWriter, *http.Request))
		mux["/qps"] = qps
		mux["/prof"] = prof
		go func() {
			err := server.ListenAndServe()

			if err != nil {
				log.Fatal(err)
			}
		}()
		log.Println("start admin ,port:", conf.Globals.AdminPort)
	}
}

type AdminHandle struct{}

func (*AdminHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
	}
	io.WriteString(w, "URL"+r.URL.String())
}

func qps(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "qps")
}

func prof(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "prof")
}
