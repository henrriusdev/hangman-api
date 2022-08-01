package main

import (
	"html/template"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./UI/base.tmpl.html",
		"./UI/pages/home.tmpl.html",
	}
	res, err := template.ParseFiles(files...)
	err = res.ExecuteTemplate(w, "base", nil)
	thereAreAnError(err)
}
