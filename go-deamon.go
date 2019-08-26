package main

import (
	"fmt"
	//	"html"
	//	"log"
	//	"net/http"

	"./rconf"
	"./worker"

	"github.com/sevlyar/go-daemon"
)

// To terminate the daemon use:
//  kill `cat sample.pid`
func main() {
	cntxt := &daemon.Context{}
	cntxt = rconf.GetDeamonConfig()
	workercfg := &worker.WConfig{}
	workercfg = rconf.GetWorkerConfig()

	fmt.Println(cntxt.LogFileName)
	fmt.Println(workercfg.WorkerSleepTime)
	/*
	   	d, err := cntxt.Reborn()
	   	if err != nil {
	   		log.Fatal("Unable to run: ", err)
	   	}
	   	if d != nil {
	   		return
	   	}
	   	defer cntxt.Release()

	   	log.Print("- - - - - - - - - - - - - - -")
	   	log.Print("daemon started")

	   	serveHTTP()
	   }

	   func serveHTTP() {
	   	http.HandleFunc("/", httpHandler)
	   	http.ListenAndServe("127.0.0.1:8080", nil)
	   }

	   func httpHandler(w http.ResponseWriter, r *http.Request) {
	   	log.Printf("request from %s: %s %q", r.RemoteAddr, r.Method, r.URL)
	   	fmt.Fprintf(w, "go-daemon: %q", html.EscapeString(r.URL.Path))*/
}
