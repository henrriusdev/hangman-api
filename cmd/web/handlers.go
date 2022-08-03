package main

import (
	"encoding/json"
	"fmt"
	"gamesapi/hangman/internal/models"
	"log"
	"net/http"
	"strconv"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func getHangman(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if r.Method != http.MethodGet {
		fmt.Fprintf(w, "Method not allowed. %d", http.StatusMethodNotAllowed)
		return
	}
	if r.URL.Path == "/new/hangman" {
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Fatal(err)
	}

	quest := models.GetQuestion(id)

	json.NewEncoder(w).Encode(quest)
}

func createHangman(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	hint, answer := r.URL.Query().Get("hint"), r.URL.Query().Get("answer")
	fmt.Println(hint, answer)
	err := models.Insert(hint, answer)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("Added hint: %s to the database!", hint)
}
