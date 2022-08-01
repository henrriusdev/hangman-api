package main

import (
	"log"
	"net/http"
)

func thereAreAnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	err := http.ListenAndServe(":4000", mux)
	thereAreAnError(err)
}
