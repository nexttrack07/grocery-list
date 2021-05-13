package main

import (
	"html/template"
	"log"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/base.layout.tmpl")
	if err != nil {
		log.Fatal("template not found", err)
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Fatal("Unable to execute template", err)
	}
}
