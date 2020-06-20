package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type authentication struct {
	authenticated bool
	username      string
}

var tmplt *template.Template

func init() {
	// ParseGlob will find all templates matching the pattern and store the templates into tmplt type.
	// Which is again of *templete.Template type
	parsedTemplates, err := template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatal("Something went wrong while parsing html files!", err)
	}

	// Must functions, used to verify that a template is valid during parsing.
	tmplt = template.Must(parsedTemplates, err)
}

func loginHandler(writer http.ResponseWriter, request *http.Request) {

	fmt.Println("Method:", request.Method)

	if request.Method == "GET" {
		tmplt.ExecuteTemplate(writer, "login.html", 1)

	} else {
		request.ParseForm()
		fmt.Println("Username:", request.Form["username"])
		fmt.Println("Password:", request.Form["password"])
		tmplt.ExecuteTemplate(writer, "login.html", 1)

	}
}

func SignupHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "This is for creating account!")
}

func LandingPageHandler(writer http.ResponseWriter, request *http.Request) {
	temp, err := template.ParseFiles("templates/index.html")

	if err != nil {
		log.Fatal("Not able to render login page", err)
	}
	temp.Execute(writer, nil)
}
