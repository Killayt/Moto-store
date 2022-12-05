package main

import (
	"log"
	"net/http"

	"github.com/Killayt/Moto-store/pkg/handlers"
	"github.com/gorilla/mux"
)

func handleRequest() error {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.Home)
	r.HandleFunc("/catalog", handlers.Catalog)
	r.HandleFunc("/contact", handlers.Contact)
	r.HandleFunc("/about", handlers.AboutUs)
	r.HandleFunc("/payment", handlers.Payment)

	http.Handle("/", r)

	// boot
	err := http.ListenAndServe(":8080", nil)

	return err
}

func main() {
	if err := handleRequest(); err != nil {
		log.Fatal("Error to start server")
	}
}
