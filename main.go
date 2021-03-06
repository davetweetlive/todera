package main

import (
	"fmt"
	"madhyam/logging"
	"madhyam/middleware"
	"madhyam/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// Create a mux router
	r := mux.NewRouter()
	r.HandleFunc("/", routes.HomePageGetHandler).Methods("GET")
	r.HandleFunc("/login", routes.LoginGetHandler).Methods("GET")
	r.HandleFunc("/login", routes.LoginPostHandler).Methods("POST")
	r.HandleFunc("/signup", routes.SignupGetHandler).Methods("GET")
	r.HandleFunc("/signup", routes.SignupPostHandler).Methods("POST")
	r.HandleFunc("/logout", routes.LogoutGetHandler).Methods("GET")
	r.HandleFunc("/me/{user}", middleware.AuthRequired(routes.PofileGetHandler)).Methods("GET")

	// Routes for learning purpose which will be refactored and code will be
	// shifted to their respective file
	r.HandleFunc("/index", routes.TestTemplate).Methods("GET")
	r.HandleFunc("/category", routes.CategoriesHandler).Methods("GET")

	r.HandleFunc("/makemigrations", routes.MigrationHandler).Methods("GET")

	// Static file server
	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	http.Handle("/", r)

	fmt.Println("Starting the server at port :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		logging.WriteLog(logging.FATAL, "Couldn't start the server")
	}
}
