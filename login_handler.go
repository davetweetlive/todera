package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func loginHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Method:", request.Method)
	// fmt.Fprintln(writer, "Login Handler!")
	if request.Method == "GET" {
		temp, err := template.ParseFiles("templates/login.html")
		if err != nil {
			log.Fatal("Not able to render login page", err)
		}
		temp.Execute(writer, nil)
	} else {
	}
}

func SignupHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "This is for creating account!")
}

func LandingPageHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(writer, "Home Page")
}