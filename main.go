package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Starting to listen..")
	http.Handle("/docs/", http.StripPrefix("/docs/", http.FileServer(http.Dir("./doc"))))
	log.Fatal(http.ListenAndServe(":3000", nil))
}
