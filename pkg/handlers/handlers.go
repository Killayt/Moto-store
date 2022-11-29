package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/index.html")
	if err != nil {
		fmt.Println("Error parse templates\n\n", err.Error())
	}
	tmpl.Execute(w, tmpl)

}

func AboutUs(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/about.html")
	if err != nil {
		fmt.Println("Error parse templates\n\n", err.Error())
	}
	tmpl.Execute(w, tmpl)
}

func Contact(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/contact.html")
	if err != nil {
		fmt.Println("Error parse templates\n\n", err.Error())
	}
	tmpl.Execute(w, tmpl)
}

func Payment(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/payment.html")
	if err != nil {
		fmt.Println("Error parse templates\n\n", err.Error())
	}
	tmpl.Execute(w, tmpl)
}

func Catalog(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/catalog.html")
	if err != nil {
		fmt.Println("Error parse templates\n\n", err.Error())
	}
	tmpl.Execute(w, tmpl)
}
