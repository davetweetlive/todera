package main

import (
	"Blog/views"
	"net/http"

	"github.com/gorilla/mux"
)

// func init() {
// 	models.CreateUserTable()
// }

func main() {

	// Static files server like css, JavaScript and image files
	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// Create a mux router
	r := mux.NewRouter()
	r.HandleFunc("/", views.HomePageGetHandler).Methods("GET")
	r.HandleFunc("/signup", views.SignUpGetHandler).Methods("GET")
	r.HandleFunc("/signup", views.SignUpPostHandler).Methods("POST")
	r.HandleFunc("/login", views.LoginGetHandler).Methods("GET")
	r.HandleFunc("/login", views.LoginPostHandler).Methods("POST")
	http.Handle("/", r)
	http.ListenAndServe(":8080", r)
}
