package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template
var port = ":8000"

type User struct {
	name     string
	password string
}

var users = []User{
	{
		name: "marcelo", password: "marcelo",
	},
	{
		name: "josiane", password: "josi1977",
	},
}

type Balance struct {
	Name   string
	Limit  float64
	Actual float64
}

var Balances = []Balance{
	{Name: "Açougue", Limit: 400, Actual: 250},
	{Name: "Alimentação", Limit: 700, Actual: 0},
	{Name: "Casa", Limit: 300, Actual: 0},
	{Name: "Cabeleleira/Manicure", Limit: 400, Actual: 0},
	{Name: "Investimento de carreira", Limit: 200, Actual: 0},
	{Name: "Consultório", Limit: 200, Actual: 0},
	{Name: "Desconhecido", Limit: 200, Actual: 0},
	{Name: "Diversos", Limit: 600, Actual: 0},
	{Name: "Farmácia", Limit: 1000, Actual: 0},
	{Name: "Mercado", Limit: 1600, Actual: 0},
	{Name: "Padaria", Limit: 600, Actual: 0},
	{Name: "Papelaria", Limit: 200, Actual: 0},
	{Name: "Perfumaria", Limit: 200, Actual: 0},
	{Name: "Presente", Limit: 300, Actual: 0},
	{Name: "Roupas", Limit: 500, Actual: 0},
	{Name: "Transporte", Limit: 600, Actual: 0},
}

func home(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "home.html", nil)
}

func checkCredentials(name string, password string) bool {
	for _, value := range users {
		if value.name == name {
			if value.password == password {
				return true
			}
		}
	}
	fmt.Println("Usuario invalido")
	return false
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

	templates.ExecuteTemplate(w, "balance.html", Balances)
}

func setupRoutes() {
	http.HandleFunc("/", home)
	http.HandleFunc("/login", login)
	http.HandleFunc("/invaliduser", invalidUser)
	http.HandleFunc("/balance", balance)
}

func main() {
	templates = template.Must(template.ParseGlob("./web/*.html"))

	setupRoutes()

	fmt.Printf("Listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
