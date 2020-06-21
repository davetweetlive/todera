package main

import (
	"Blog/models"
	"log"
	"net/http"
)

// var SQLConnector *sql.DB

// func init() {
// 	// Connect to the database
// 	SQLConnector, _ = models.EstablishDBConnection()
// }

func main() {
	models.CreateTable()
	// fmt.Println(SQLConnector)
	// To serve static fils
	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	// different handlers
	http.HandleFunc("/", LandingPageHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/signup", SignupHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln("")
	}
}
