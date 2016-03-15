package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	err := http.ListenAndServeTLS(":3000", "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
