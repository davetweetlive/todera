package models

import (
	"fmt"
)

type User struct {
	username string
	email    string
	password string
}

// var db *sql.DB = ConDB()
func NewUser(username, email, password string) {
	// Check if the username is taken
	sqlStatement := "select username from users where username = ?"

	var user string
	db := ConDB()
	defer db.Close()

	row := db.QueryRow(sqlStatement, username)

	if err := row.Scan(&user); err == nil {
		fmt.Println("The username is already taken")
		fmt.Println(user)
	}

	// Check if the email is already registered

	// Return new user

}

func RegisterUser(username, email, password string) error {

	// Create a new user
	// cost := bcrypt.DefaultCost
	// hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	// if err != nil {
	// 	fmt.Println("Error occoured while generating hash")
	// }

	// row, err := db.Query("")
	// return nil
	return nil
}
