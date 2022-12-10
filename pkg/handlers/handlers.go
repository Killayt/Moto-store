package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/index.html", "web/templates/header.html", "web/templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, "Error with `index` template")
	}
	tmpl.ExecuteTemplate(w, "index", nil)
}

func AboutUs(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/about.html", "web/templates/header.html", "web/templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, "Error with `about us` template")
	}
	tmpl.ExecuteTemplate(w, "about", nil)
}

func Contact(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/header.html", "web/templates/contact.html", "web/templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, "Error with `contact` template")
	}
	tmpl.ExecuteTemplate(w, "contact", nil)
}

func Payment(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/payment.html", "web/templates/header.html", "web/templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, "Error with `payment` template")
	}
	tmpl.ExecuteTemplate(w, "payment", nil)
}

func Catalog(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/catalog.html", "web/templates/header.html", "web/templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, "Error with `catalog` template")
	}
	tmpl.ExecuteTemplate(w, "catalog", nil)
}

// Methods

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/methods/createProduct.html", "web/templates/header.html", "web/templates/footer.html")
	if err != nil {
		fmt.Fprintf(w, "Erorr with template. Try again")
	}
	tmpl.ExecuteTemplate(w, "createProduct", nil)
}
