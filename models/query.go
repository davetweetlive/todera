package models

const (
	sqlStatementForUsername = "SELECT username FROM users WHERE username = ?"
	sqlStatementForEmail    = "SELECT email FROM users WHERE email = ?"

	userCreate = "INSERT INTO users (username, email, password, created_at) VALUES (?, ?, ?, ?)"
)
