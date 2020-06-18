package main

import (
	"log"
	"net/http"
)

func main() {
	// To serve static fils
	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	http.HandleFunc("/", LandingPageHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/signup", SignupHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln("")
	}
}
