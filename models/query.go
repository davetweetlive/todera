package models

const (
	selectUsername = "SELECT username FROM users WHERE username = ?"
	selectEmail    = "SELECT email FROM users WHERE email = ?"
	createUser     = "INSERT INTO users (username, password, email, created_at) VALUES (?, ?, ?, ?);"
)

const (
	createBlogTable = ` `
)
