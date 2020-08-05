package main

import (
	"fmt"
	"madhyam/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// Create a mux router
	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomePageGetHandler).Methods("GET")
	r.HandleFunc("/login", routes.LoginGetHandler).Methods("GET")

	// Static file server
	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	http.Handle("/", r)

	fmt.Println("Starting the server at port :8080")
	http.ListenAndServe(":8080", r)
}
