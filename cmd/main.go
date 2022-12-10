package main

import (
	"log"
	"net/http"

	"github.com/Killayt/Moto-store/pkg/handlers"
	"github.com/gorilla/mux"
)

func handleRequest() error {
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.Index).Methods("GET")
	r.HandleFunc("/createProduct", handlers.CreateProduct).Methods("GET")
	r.HandleFunc("/saveArticle", handlers.SaveArticle).Methods("POST")
	r.HandleFunc("/catalog", handlers.Catalog).Methods("GET")
	r.HandleFunc("/contact", handlers.Contact).Methods("GET")
	r.HandleFunc("/about", handlers.AboutUs).Methods("GET")
	r.HandleFunc("/payment", handlers.Payment).Methods("POST")

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
