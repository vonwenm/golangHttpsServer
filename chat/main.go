// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"net/http"
	"text/template"
	"fmt"
)

const (
    PORT       = ":8443"
    PRIV_KEY   = "key.pem"
    PUBLIC_KEY = "cert.pem"
)

var addr = flag.String("addr", ":8787", "http service address")
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
	//err := http.ListenAndServe(*addr, nil)
	//if err != nil {
	//	log.Fatal("ListenAndServe: ", err)
	//}
	err := http.ListenAndServeTLS(PORT, PUBLIC_KEY, PRIV_KEY, nil)
	if err != nil {
		fmt.Printf("main(): %s\n", err)
	        log.Fatal(err)
    	}  
}
