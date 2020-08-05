package routes

import (
	"database/sql"
	"html/template"
	"madhyam/models"
	"net/http"
	"time"
)

var templates *template.Template

type PageInfo struct {
	Title      string
	Owner      string
	LastUpdate time.Time
}

var db *sql.DB

func init() {
	templates = template.Must(template.ParseGlob("templates/*.html"))
	// Database
	db = models.ConDB()
}

func HomePageGetHandler(w http.ResponseWriter, r *http.Request) {
	t := PageInfo{Title: "Madhyam", Owner: "SAffron Coders", LastUpdate: time.Now()}
	templates.ExecuteTemplate(w, "index.html", t)
}

func LoginGetHandler(w http.ResponseWriter, r *http.Request) {
	t := PageInfo{Title: "Madhyam", Owner: "SAffron Coders", LastUpdate: time.Now()}
	templates.ExecuteTemplate(w, "login.html", t)
}

func LoginPostHandler(w http.ResponseWriter, r *http.Request) {

}

func SignupGetHandler(w http.ResponseWriter, r *http.Request) {
	t := PageInfo{Title: "Madhyam", Owner: "SAffron Coders", LastUpdate: time.Now()}
	templates.ExecuteTemplate(w, "signup.html", t)
}
func SignupPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")
	password1 := r.PostForm.Get("password1")
	if password != password1 {
		templates.ExecuteTemplate(w, "signup.html", "Password didn't match")
	}
	models.RegisterUser(username, email, password)

	http.Redirect(w, r, "/login", http.StatusFound)

}
