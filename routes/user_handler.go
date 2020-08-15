package routes

import (
	"database/sql"
	"fmt"
	"html/template"
	"madhyam/info"
	"madhyam/models"
	"madhyam/sessions"
	"madhyam/utils"
	"net/http"

	"github.com/gorilla/mux"
)

var templates *template.Template
var db *sql.DB

// Initialization of variables
var pageInfo *info.PageInfo

func init() {
	templates = template.Must(template.ParseGlob("templates/*.html"))
	db = models.ConDB()
	pageInfo = info.GetPageInfo()
}

// I
// N
// D
// E
// X
// Page handler if authenticated then users details will be displayed
// Else just usual information will be rendered including login signup
// options.
func HomePageGetHandler(w http.ResponseWriter, r *http.Request) {
	session, err := sessions.Store.Get(r, "session")
	fmt.Println("Session value", session)
	if err != nil {
		fmt.Println("Session not available")
		templates.ExecuteTemplate(w, "index.html", pageInfo)
		return
	}
	if session == nil {
		pageInfo.IsAuthenticated = false
		pageInfo.AuthenticationDetails.SessionVal = nil
		fmt.Println("Not authenticated:", pageInfo.AuthenticationDetails.SessionVal)
		templates.ExecuteTemplate(w, "index.html", pageInfo)
		return
	}

	fmt.Println("Logged in as:", pageInfo.AuthenticationDetails.SessionVal)
	templates.ExecuteTemplate(w, "index.html", pageInfo)

}

// L
// O
// G
// I
// N
// Handler for GET method
func LoginGetHandler(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "login.html", pageInfo); err != nil {
		fmt.Println("Couldn't find login template")
	}
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
	session.Values["username"] = username
	session.Save(r, w)

	pageInfo.IsAuthenticated = true
	pageInfo.AuthenticationDetails.SessionVal = session.Values["username"]
	// Removed StatusFound which is 302 and added StatusSeeOther which is 303
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// S
// I
// G
// N
// Up GET handler
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
			pageInfo.ErrorMsg = models.ErrUsernameIsTaken
			templates.ExecuteTemplate(w, "signup.html", pageInfo)

		case models.ErrEmailAlreadyRegistered:

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
	session.Options.MaxAge = -1
	delete(session.Values, "session")
	pageInfo.IsAuthenticated = false
	if err := session.Save(r, w); err != nil {
		fmt.Println("Couldn't delete the session")
	}
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// P
// R
// O
// F
// I
// L
// E
func PofileGetHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userId := params["user"]
	fmt.Println(userId)

	templates.ExecuteTemplate(w, "profile.html", pageInfo)
}
