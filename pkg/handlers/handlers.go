package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/Killayt/Moto-store/pkg/model"
)

func Home(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/index.html")
	if err != nil {
		fmt.Println("Error parse templates\n\n", err.Error())
	}

	homePage := model.Page{
		Title:   "Home",
		Text:    "hueta",
		Product: model.Product{},
	}

	tmpl.Execute(w, homePage)
}

func AboutUs(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/about.html")
	if err != nil {
		fmt.Println("Error parse templates\n\n", err.Error())
	}

	aboutPage := model.Page{
		Title:   "About us",
		Text:    "hueta",
		Product: model.Product{},
	}

	tmpl.Execute(w, aboutPage)
}

func Contact(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/contact.html")
	if err != nil {
		fmt.Println("Error parse templates\n\n", err.Error())
	}

	contactPage := model.Page{
		Title:   "Contact",
		Text:    "hueta",
		Product: model.Product{},
	}

	tmpl.Execute(w, contactPage)
}

func Payment(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/payment.html")
	if err != nil {
		fmt.Println("Error parse templates\n\n", err.Error())
	}

	paymentPage := model.Page{
		Title:   "Payment",
		Text:    "hueta",
		Product: model.Product{},
	}

	tmpl.Execute(w, paymentPage)
}

func Catalog(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/catalog.html")
	if err != nil {
		fmt.Println("Error parse templates\n\n", err.Error())
	}

	catalogPage := model.Page{
		Title:   "Catalog",
		Text:    "hueta",
		Product: model.Product{},
	}

	tmpl.Execute(w, catalogPage)
}
