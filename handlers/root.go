package handlers

import (
	"html/template"
	"log"
	"net/http"

	assets "github.com/rakicjovan/WebWake/templates"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFS(assets.Assets, "index.html")
	if err != nil {
		log.Printf("Error parsing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Printf("Error executing template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
