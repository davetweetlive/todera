package models

import (
	"database/sql"
	"errors"
	"fmt"
	"madhyam/models/query"
	"madhyam/utils"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

type User struct {
	username     string
	password     []byte
	email        string
	firstName    string
	lastName     string
	createdAt    time.Time
	isASuperUser bool
	ProfileImage []byte
}

var (
	ErrEmailAlreadyRegistered   = errors.New("The email is already registered")
	ErrUsernameIsTaken          = errors.New("The username is already taken")
	ErrUserNotAvailable         = errors.New("User isn't available")
	ErrUsernamePasswordMismatch = errors.New("Username or password wrong")
)

var uname string

func init() {
	db = ConDB()
}

// To check whether the email provided by a user is already registered or not
func (u *User) ifEmailRegistered() bool {
	fmt.Println("Email already registered mothod")
	if err := db.QueryRow(query.IfEmailExists, u.email).Scan(&uname); err == nil {
		return true
	}
	return false
}

// To check whether the user name has been already registered or taken by other user of the syatem
func (u *User) ifUsernameTaken() bool {
	if err := db.QueryRow(query.IfUsernameExists, u.username).Scan(&uname); err == nil {
		return true
	}
	return false
}

// Temproarily written and It will be refactored shortly
func NewUser(username, email string, hash []byte) (*User, error) {
	return &User{
		username:     username,
		email:        email,
		password:     hash,
		isASuperUser: false,
	}, nil
}

// Registers a new user to the system
func RegisterUser(username, email, password string) error {

	// 1. Get passwords hash
	hash, _ := utils.GetHash(password)

	//2. Get the new user
	user, err := NewUser(username, email, hash)
	if err != nil {
		fmt.Println("Error while creating a new user!")
		return err
	}

	// 3. Check whether the email and username is already registered
	if user.ifEmailRegistered() {
		fmt.Println("Email already registered")
		return ErrEmailAlreadyRegistered

	} else if user.ifUsernameTaken() {
		fmt.Println("Username already taken")
		return ErrUsernameIsTaken
	}

	// 5. Register user
	_, err = db.Query(query.CreateUser, username, string(hash), email, false)
	if err != nil {
		fmt.Println("Couldn't register!")
		return err
	}

	return nil
}

// Steps for authentication an user via username, password or email, password

func AuthenticateUser(username, password string) error {
	var hash string

	if err := db.QueryRow("SELECT  `password` FROM user where username = ? or email = ?;", username, username).Scan(&hash); err != nil {
		fmt.Println("Login validation failed!")
		return ErrUserNotAvailable
	}
	fmt.Println(hash)

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return ErrUsernamePasswordMismatch
	}
	return err
}
