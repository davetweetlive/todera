package models

import (
	"database/sql"
	"fmt"
	"madhyam/utils"
	"time"
)

var db *sql.DB

type User struct {
	username string
	email    string
	password []byte
}

func init() {
	db = ConDB()

}

// To check whether the email provided by a user is already registered or not
func (u *User) ifEmailRegistered() bool {
	var uname string
	if err := db.QueryRow(selectEmail, u.username).Scan(&uname); err == nil {
		return true
	}
	return false

}

// To check whether the user name has been already registered or taken by other user of the syatem
func (u *User) ifUsernameTaken() bool {
	var uname string
	if err := db.QueryRow(selectUsername, u.username).Scan(&uname); err == nil {
		return true
	}
	return false
}

// Temproarily written and It will be refactored shortly
func NewUser(username, email string, hash []byte) (*User, error) {
	return &User{username, email, hash}, nil
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

	// 3. Check whether the email is already registered
	if user.ifEmailRegistered() {
		fmt.Println("Email already registered")
		return err
	}
	// 4. Check wheather the username is already used
	if user.ifUsernameTaken() {
		fmt.Println("Username already taken")
		return err
	}

	// 5. Register user
	_, err = db.Query(createUser, username, string(hash), email, time.Now())
	if err != nil {
		fmt.Println("Couldn't register!")
		return err
	}

	return nil
}

func AuthenticateUser(username, password string) error {

	return nil
}
