package main

import (
	"net/http"

	"github.com/Killayt/Moto-store/pkg/handlers"
	"github.com/gorilla/mux"
)

func handleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.Home)
	r.HandleFunc("/catalog", handlers.Catalog)
	r.HandleFunc("/contact", handlers.Contact)
	r.HandleFunc("/about", handlers.AboutUs)

	http.Handle("/", r)

	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()
}
