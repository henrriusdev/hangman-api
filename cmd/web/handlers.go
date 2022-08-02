package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		fmt.Fprintf(w, "Method not allowed. %d", http.StatusMethodNotAllowed)
		return
	}

	files := []string{
		"./UI/base.tmpl.html",
		"./UI/pages/home.tmpl.html",
	}
	res, err := template.ParseFiles(files...)
	err = res.ExecuteTemplate(w, "base", nil)
	thereAreAnError(err)
}
