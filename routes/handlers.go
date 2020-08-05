package routes

import (
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type PageInfo struct {
	Title      string
	Owner      string
	LastUpdate time.Time
}

func HomePageGetHandler(w http.ResponseWriter, r *http.Request) {
	t := PageInfo{Title: "Madhyam", Owner: "SAffron Coders", LastUpdate: time.Now()}
	templt, err := template.ParseFiles("templates/index.html")
	if err != nil {
		fmt.Println("Couldn't parse the temlate")
	}
	templt.Execute(w, t)
}

func LoginGetHandler(w http.ResponseWriter, r *http.Request) {
	t := PageInfo{Title: "Madhyam", Owner: "SAffron Coders", LastUpdate: time.Now()}
	templt, err := template.ParseFiles("templates/login.html")
	if err != nil {
		fmt.Println("Couldn't parse the temlate")
	}
	templt.Execute(w, t)
}

func LoginPostHandler(w http.ResponseWriter, r *http.Request) {

}

func SignupGetHandler(w http.ResponseWriter, r *http.Request) {

}
func SignupPostHandler(w http.ResponseWriter, r *http.Request) {

}
