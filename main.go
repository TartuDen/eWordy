package main

import (
	"eWordy/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.GetPage)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("Something is wrong - " + err.Error())
	}
}
