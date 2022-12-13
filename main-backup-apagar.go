package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template
var port = ":8080"

var title = "Saldo do mês: Outubro"
var updatedAt = "08/11/2022"

type pageContent struct {
	Title     string
	Balances  []balanceOut
	Total     balanceOut
	UpdatedAt string
}

func home(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "home.html", nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		name := r.FormValue("name")
		password := r.FormValue("password")

		validUser := checkCredentials(name, password)

		if validUser {
			balance(w, r)
		} else {
			invalidUser(w, r)
		}

	default:
		fmt.Fprintf(w, "Sorry, only POST method is supported.")
	}
}

func invalidUser(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "invaliduser.html", nil)
}

func balance(w http.ResponseWriter, r *http.Request) {

	balancesOut := getBalances()
	totalOut := getTotal()

	content := pageContent{
		Title:     title,
		Balances:  balancesOut,
		Total:     totalOut,
		UpdatedAt: updatedAt,
	}

	templates.ExecuteTemplate(w, "balance.html", content)
}

func setupRoutes() {
	http.HandleFunc("/", home)
	http.HandleFunc("/login", login)
	http.HandleFunc("/invaliduser", invalidUser)
	http.HandleFunc("/balance", balance)
}

func _main() {
	templates = template.Must(template.ParseGlob("./web/*.html"))

	setupRoutes()

	fmt.Printf("Listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
