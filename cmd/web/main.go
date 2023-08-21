package main

import (
	"net/http"
	"exercise1/pkg/handlers"
)

func main() { 
	
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	http.ListenAndServe(":2000", nil)
}