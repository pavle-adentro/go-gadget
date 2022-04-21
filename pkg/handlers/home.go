package handlers

import (
	"net/http"
	"path"
	"text/template"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, err); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
