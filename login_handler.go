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

func loginHandler(writer http.ResponseWriter, request *http.Request) {

	fmt.Println("Method:", request.Method)
	temp, err := template.ParseFiles("templates/login.html")
	if err != nil {
		log.Fatal("Not able to render login page", err)
	}

	if request.Method == "GET" {
		temp.Execute(writer, nil)

	} else {
		request.ParseForm()
		fmt.Println("Username:", request.Form["username"])
		fmt.Println("Password:", request.Form["password"])
		temp.Execute(writer, nil)

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
