package handlers

import (
	"html/template"
	"net/http"
)

type Data struct {
	EnglWord    string
	RusTrue     string
	RusVariants []string
}

func GetPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	data := Data{}

	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		http.Error(w, "Template not found - "+err.Error(), http.StatusNotFound)
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error"+err.Error(), http.StatusInternalServerError)
	}

}
