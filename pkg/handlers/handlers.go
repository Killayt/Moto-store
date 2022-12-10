package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"os"
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

func SaveArticle(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	anons := r.FormValue("anons")
	text := r.FormValue("full_text")

	if title == "" || anons == "" || text == "" { // checking validaty
		fmt.Fprintf(w, "Input form is null")
	} else {
		db, err := sql.Open("mysql", os.Getenv("DB_NAME")+":"+os.Getenv("DB_PASS")+"@tcp(127.0.0.1"+os.Getenv("DB_PORT")+")/articles")
		if err != nil {
			fmt.Fprintf(w, "Error to connect to database")
		}
		defer db.Close()

		insert, err := db.Query(fmt.Sprintf("INSERT INTO `Articles` (`title`, `anons`, `full_text`) VALUES (`%s`, `%s`, `%s`)", title, anons, text))
		if err != nil {
			fmt.Fprintf(w, "Error to insert into database")
		}
		defer insert.Close()

		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}
