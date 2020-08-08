package models

const (
	selectUsername = "SELECT username FROM users WHERE username = ?"
	selectEmail    = "SELECT email FROM users WHERE email = ?"
	createUser     = "INSERT INTO users (username, password, email, created_at) VALUES (?, ?, ?, ?);"
)

const (
	createBlogTable = `CREATE TABLE article (
		id int NOT NULL,
		content int NOT NULL,
		author int,
		PRIMARY KEY (id),
		FOREIGN KEY (author) REFERENCES user(id)
	);`

	userProfile = `CREATE TABLE profile`
)
