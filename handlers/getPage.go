package handlers

import (
	"html/template"
	"net/http"
)

func getPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		http.Error(w, "Template not found - "+err.Error(), http.StatusNotFound)
	}

	err = tmpl.Execute(w, data)

}
