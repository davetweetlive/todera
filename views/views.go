package views

import (
	"Blog/models"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var tmplt *template.Template
var client *sql.DB

type User struct {
	ID       int
	Username int
	Email    string
	Password []byte
}

func init() {
	// ParseGlob will find all templates matching the pattern and store the templates into tmplt type.
	// Which is again of *templete.Template type
	parsedTemplates, err := template.ParseGlob("templates/*.html")
	if err != nil {
		log.Fatal("Something went wrong while parsing html files!", err)
	}

	// Must functions, used to verify that a template is valid during parsing.
	tmplt = template.Must(parsedTemplates, err)

	// instantiate client database connection
	client, err = models.ConnectToMySQL()
}

func HomePageGetHandler(writer http.ResponseWriter, request *http.Request) {
	tmplt.ExecuteTemplate(writer, "index.html", 1)
}

func SignUpGetHandler(writer http.ResponseWriter, request *http.Request) {
	tmplt.ExecuteTemplate(writer, "signup.html", ' ')
}

func SignUpPostHandler(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	username := request.PostForm.Get("username")
	emailID := request.PostForm.Get("email-id")
	password := request.PostForm.Get("password")
	createdAt := time.Now()

	// Bcrypt
	cost := bcrypt.DefaultCost
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return
	}
	result, err := client.Exec(`INSERT INTO users (username, password, email, created_at) VALUES (?, ?, ?, ?)`, username, hash, emailID, createdAt)
	if err != nil {
		log.Fatalln("Couldn't sign in")
	}
	id, err := result.LastInsertId()
	fmt.Println(id)
	http.Redirect(writer, request, "/login", 200)
}

func LoginGetHandler(writer http.ResponseWriter, request *http.Request) {
	tmplt.ExecuteTemplate(writer, "login.html", ' ')
}

func LoginPostHandler(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	username := request.PostForm.Get("username")
	password := request.PostForm.Get("password")
	var user User
	result := client.QueryRow("select password from users where username=$1", username)

	err := result.Scan(&user.Password)

	fmt.Println(&user.Password)
	if err != nil {

		// If an entry with the username does not exist, send an "Unauthorized"(401) status
		if err == sql.ErrNoRows {
			fmt.Println(6)

			writer.WriteHeader(http.StatusUnauthorized)
			return
		}
		fmt.Println(7)
		// If the error is of any other type, send a 500 status
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	// fmt.Println("THe hash is ", pass)
	// fmt.Println(reflect.TypeOf(result))
	// Compare password
	if err = bcrypt.CompareHashAndPassword(user.Password, []byte(password)); err != nil {
		// If the two passwords don't match, return a 401 status
		writer.WriteHeader(http.StatusUnauthorized)
	}
	fmt.Println("You are ologged in")
}

// db.QueryRow("select password from users where username=$1", creds.Username)
