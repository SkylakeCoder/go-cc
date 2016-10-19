package chess

import (
	"net/http"
	"log"
	"io/ioutil"
)

const (
	START_COUNT = 10
)
var requestCount = 0

func StartServer() {
	http.HandleFunc("/", handleRequest)
	log.Fatalln(http.ListenAndServe("localhost:8686", nil))
}

func handleRequest(w http.ResponseWriter, req *http.Request) {
	requestCount++
	chess := req.FormValue("chess")
	needMaster := false
	if requestCount <= START_COUNT {
		resp, err := http.Get("http://localhost:8688?chess=" + chess)
		if err != nil {
			log.Fatalln("error when communicate with repository server...")
		}
		defer resp.Body.Close()
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln("error when read body...")
		}
		if len(bytes) != 0 {
			w.Write(bytes)
		} else {
			needMaster = true
		}
	} else {
		needMaster = true
	}
	if needMaster {
		//@TODO...
	}
}
