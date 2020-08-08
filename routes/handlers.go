package routes

import (
	"database/sql"
	"fmt"
	"html/template"
	"madhyam/models"
	"madhyam/sessions"
	"madhyam/utils"
	"net/http"
	"time"
)

type PageInfo struct {
	Title      string
	Owner      string
	LastUpdate time.Time
	ErrorMsg   interface{}
}

var templates *template.Template
var db *sql.DB
var pageInfo PageInfo

// Initialization of variables

func init() {
	templates = template.Must(template.ParseGlob("templates/*.html"))
	// Database
	db = models.ConDB()

	// Page info passed to the template from the backend
	pageInfo = PageInfo{
		Title:      "Madhyam",
		Owner:      "Saffron Coders",
		LastUpdate: time.Now(),
		ErrorMsg:   nil,
	}
}

// I
// N
// D
// E
// X
// Page handler
func HomePageGetHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", pageInfo)
}

func LoginGetHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "login.html", pageInfo)
}

// L
// O
// G
// I
// N
// Handler for POST method
func LoginPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	password := r.PostForm.Get("password")

	if err := models.AuthenticateUser(username, password); err != nil {

		if err == models.ErrUserNotAvailable {
			pageInfo.ErrorMsg = models.ErrUserNotAvailable
			templates.ExecuteTemplate(w, "login.html", pageInfo)

		} else if err == models.ErrUsernamePasswordMismatch {
			pageInfo.ErrorMsg = models.ErrUsernamePasswordMismatch
			templates.ExecuteTemplate(w, "login.html", pageInfo)

		} else {
			utils.InternalServerError(w)
		}
	}

	// Session
	session, err := sessions.Store.Get(r, "session")
	if err != nil {
		fmt.Println("Can't find the session")
	}
	session.Values["userid"] = username
	session.Save(r, w)

	// Removed StatusFound which is 302 and added StatusSeeOther which is 303
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func SignupGetHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "signup.html", pageInfo)
}

// S
// I
// G
// N
// Up constrains for POST method
func SignupPostHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.PostForm.Get("username")
	email := r.PostForm.Get("email")
	password := r.PostForm.Get("password")
	password1 := r.PostForm.Get("password1")

	if password != password1 {
		templates.ExecuteTemplate(w, "signup.html", "Password didn't match")
	}
	if err := models.RegisterUser(username, email, password); err != nil {
		switch err {

		case models.ErrUsernameIsTaken:
			fmt.Println("Handler: Case usernam is taken")
			pageInfo.ErrorMsg = models.ErrUsernameIsTaken
			templates.ExecuteTemplate(w, "signup.html", pageInfo)

		case models.ErrEmailAlreadyRegistered:
			fmt.Println("Handler: Email is already registered")

			pageInfo.ErrorMsg = models.ErrEmailAlreadyRegistered
			templates.ExecuteTemplate(w, "signup.html", pageInfo)

		default:
			utils.InternalServerError(w)
		}
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// L
// O
// G
// O
// U
// T
// Handler for terminationg session
func LogoutGetHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := sessions.Store.Get(r, "session")

	delete(session.Values, "userid")
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
