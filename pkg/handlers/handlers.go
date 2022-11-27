package handlers

import (
	"fmt"
	"net/http"

	"github.com/Killayt/Moto-store/pkg/model"
)

func Home(w http.ResponseWriter, r *http.Request) {
	var bob model.Product
	bob1 := bob.NewProduct(1, "Bob", "BOB IS BOB", 14.99, 0)
	fmt.Fprintf(w, bob1.Name)
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
