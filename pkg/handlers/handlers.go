package handlers

import (
	"html/template"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/index.html", "web/templates/header.html", "web/templates/footer.html")
	if err != nil {
		log.Errorln("Error parse templates\n\n", err.Error())
	}
	tmpl.ExecuteTemplate(w, "index", nil)
}

func AboutUs(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/about.html", "web/templates/header.html", "web/templates/footer.html")
	if err != nil {
		log.Errorln("Error parse templates\n\n", err.Error())
	}
	tmpl.ExecuteTemplate(w, "about", nil)
}

func Contact(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/header.html", "web/templates/contact.html", "web/templates/footer.html")
	if err != nil {
		log.Errorln("Error parse templates\n\n", err.Error())
	}
	tmpl.ExecuteTemplate(w, "contact", nil)
}

func Payment(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/payment.html", "web/templates/header.html", "web/templates/footer.html")
	if err != nil {
		log.Errorln("Error parse templates\n\n", err.Error())
	}
	tmpl.ExecuteTemplate(w, "payment", nil)
}

func Catalog(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/templates/catalog.html", "web/templates/header.html", "web/templates/footer.html")
	if err != nil {
		log.Errorln("Error parse templates\n\n", err.Error())
	}
	tmpl.ExecuteTemplate(w, "catalog", nil)
}
