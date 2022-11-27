package handlers

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home")
}

func AboutUs(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "AboutUs")
}

func Contact(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contact")
}

func SendAndPay(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "SendAndPay")
}

func Catalog(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Catalog")
}
