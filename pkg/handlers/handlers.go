package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/Killayt/Moto-store/pkg/model"
)

func Home(w http.ResponseWriter, r *http.Request) {
	homePage := model.Page{
		Title:   "Home",
		Text:    "This is home page",
		Product: model.Product{},
	}

	tmpl, err := template.ParseFiles("web/templates/index.html", "web/templates/header.html", "web/templates/footer.html")
	if err != nil {
		fmt.Println("Error parse templates\n\n", err.Error())
	}
	tmpl.ExecuteTemplate(w, "index", homePage)
}

func AboutUs(w http.ResponseWriter, r *http.Request) {
	aboutPage := model.Page{
		Title:   "About us",
		Text:    "hueta",
		Product: model.Product{},
	}

	tmpl, err := template.ParseFiles("web/front-end/about.html")
	if err != nil {
		fmt.Println("Error parse templates\n\n", err.Error())
	}
	tmpl.Execute(w, aboutPage)
}

func Contact(w http.ResponseWriter, r *http.Request) {
	contactPage := model.Page{
		Title:   "Contact",
		Text:    "hueta",
		Product: model.Product{},
	}

	tmpl, err := template.ParseFiles("web/front-end/contact.html")
	if err != nil {
		fmt.Println("Error parse templates\n\n", err.Error())
	}
	tmpl.Execute(w, contactPage)
}

func Payment(w http.ResponseWriter, r *http.Request) {
	paymentPage := model.Page{
		Title:   "Payment",
		Text:    "hueta",
		Product: model.Product{},
	}

	tmpl, err := template.ParseFiles("web/front-end/payment.html")
	if err != nil {
		fmt.Println("Error parse templates\n\n", err.Error())
	}
	tmpl.Execute(w, paymentPage)
}

func Catalog(w http.ResponseWriter, r *http.Request) {
	catalogPage := model.Page{
		Title:   "Catalog",
		Text:    "hueta",
		Product: model.Product{},
	}

	tmpl, err := template.ParseFiles("web/fornd-end/catalog.html")
	if err != nil {
		fmt.Println("Error parse templates\n\n", err.Error())
	}
	tmpl.Execute(w, catalogPage)
}
