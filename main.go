package main

import (
	"net/http"
	"ewordy/handlers"
)

type Data struct {
	EnglWord    string
	RusTrue     string
	RusVariants []string
}

func main() {
	http.HandleFunc("/", handlers.getPage)
}
