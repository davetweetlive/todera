package models

import (
	"database/sql"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

type User struct {
	username string
	email    string
	password []byte
}

func init() {
	db = ConDB()
	// defer db.Close()
}

func NewUser(username, email string, hash []byte) (*User, error) {

	var queryStr string

	// Check if the username is taken
	row := db.QueryRow(sqlStatementForUsername, username)
	if err := row.Scan(&queryStr); err == nil {
		fmt.Println("The username is already taken")
		fmt.Println(queryStr)
		return nil, err
	}
	// Check if the email is already registered
	row = db.QueryRow(sqlStatementForEmail, email)
	if err := row.Scan(&queryStr); err == nil {
		fmt.Println("The email is already registered!")
		fmt.Println(queryStr)
		return nil, err
	}

	return &User{username, email, hash}, nil
	// Return new user

}

func RegisterUser(username, email, password string) error {

	cost := bcrypt.DefaultCost
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		fmt.Println("Error occoured while generating hash")
		return err
	}
	_, err = NewUser(username, email, hash)
	if err != nil {
		fmt.Println("CAn't create a new user!")
		return err
	}

	row, err := db.Query("INSERT INTO users (username, password, email, created_at) VALUES (?, ?, ?, ?);", username, string(hash), email, time.Now())
	if err != nil {
		fmt.Println("Couldn't register!")
		return err
	}
	fmt.Println(row)
	return nil
}
