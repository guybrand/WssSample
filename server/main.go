// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

var addr = flag.String("addr", ":9200", "http service address")
var homeTempl = template.Must(template.ParseFiles("home.html"))

func serveHome(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Not found", 404)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	homeTempl.Execute(w, r.Host)
}

func main() {
	flag.Parse()
	go h.run()
	http.HandleFunc("/", serveHome)
	http.HandleFunc("/ws", serveWs)
	//	err := http.ListenAndServe(*addr, nil)
	//	if err != nil {
	//		log.Fatal("ListenAndServe: ", err)
	//	}
	var httpErr error

	if _, err := os.Stat("./algo.crt"); err == nil {
		fmt.Println("file ", "algo.crt found switching to https")
		if httpErr = http.ListenAndServeTLS(*addr, "algo.crt", "algo.key", nil); httpErr != nil {
			log.Fatal("The process exited with https error: ", httpErr.Error())
		}
	} else {
		httpErr = http.ListenAndServe(*addr, nil)
		if httpErr != nil {
			log.Fatal("The process exited with http error: ", httpErr.Error())
		}
	}

}
