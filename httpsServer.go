package main


import (
	"fmt"
	"log"
	"net/http"	
)

const (
    PORT       = ":8443"
    PRIV_KEY   = "key.pem"
    PUBLIC_KEY = "cert.pem"
)

func rootHander(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Nobody should read this.")
}

func main() {
    http.HandleFunc("/", rootHander)
    log.Println("Server starting")  
    err := http.ListenAndServeTLS(PORT, PUBLIC_KEY, PRIV_KEY, nil)
    if err != nil {
        fmt.Printf("main(): %s\n", err)
        		log.Fatal(err)
    }
}