package main

import (
	"fmt"
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
	mux.HandleFunc("/get/hangman", getHangman)
	mux.HandleFunc("/new/hangman", createHangman)
	fmt.Println("Starting server...")
	err := http.ListenAndServe(":4000", mux)
	thereAreAnError(err)
}
