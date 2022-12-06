package handlers

import (
	"html/template"
	"net/http"

	"github.com/Killayt/Moto-store/pkg/model"
	log "github.com/sirupsen/logrus"
)

func Home(w http.ResponseWriter, r *http.Request) {
	homePage := model.Page{
		Title:   "Home",
		Text:    "This is home page",
		Product: model.Product{},
	}

	tmpl, err := template.ParseFiles("web/templates/index.html", "web/templates/header.html", "web/templates/footer.html")
	if err != nil {
		log.Errorln("Error parse templates\n\n", err.Error())
	}
	tmpl.ExecuteTemplate(w, "index", homePage)
}

func AboutUs(w http.ResponseWriter, r *http.Request) {
	aboutPage := model.Page{
		Title:   "About us",
		Text:    "hueta",
		Product: model.Product{},
	}

	tmpl, err := template.ParseFiles("web/templates/about.html", "web/templates/header.html", "web/templates/footer.html")
	if err != nil {
		log.Errorln("Error parse templates\n\n", err.Error())
	}
	tmpl.ExecuteTemplate(w, "about", aboutPage)
}

func Contact(w http.ResponseWriter, r *http.Request) {
	contactPage := model.Page{
		Title:   "Contact",
		Text:    "hueta",
		Product: model.Product{},
	}

	tmpl, err := template.ParseFiles("web/templates/header.html", "web/templates/contact.html", "web/templates/footer.html")
	if err != nil {
		log.Errorln("Error parse templates\n\n", err.Error())
	}
	tmpl.ExecuteTemplate(w, "contact", contactPage)
}

func Payment(w http.ResponseWriter, r *http.Request) {
	paymentPage := model.Page{
		Title:   "Payment",
		Text:    "hueta",
		Product: model.Product{},
	}

	tmpl, err := template.ParseFiles("web/templates/payment.html", "web/templates/header.html", "web/templates/footer.html")
	if err != nil {
		log.Errorln("Error parse templates\n\n", err.Error())
	}
	tmpl.ExecuteTemplate(w, "payment", paymentPage)
}

func Catalog(w http.ResponseWriter, r *http.Request) {
	catalogPage := model.Page{
		Title:   "Catalog",
		Text:    "hueta",
		Product: model.Product{},
	}

	tmpl, err := template.ParseFiles("web/templates/catalog.html", "web/templates/header.html", "web/templates/footer.html")
	if err != nil {
		log.Errorln("Error parse templates\n\n", err.Error())
	}
	tmpl.ExecuteTemplate(w, "catalog", catalogPage)
}
