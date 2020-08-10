package models

const (
	selectUsername = "SELECT username FROM users WHERE username = ?"
	selectEmail    = "SELECT email FROM users WHERE email = ?"
	createUser     = "INSERT INTO users (username, password, email, created_at) VALUES (?, ?, ?, ?);"
)

const (
	// Table for article where id (Primary key)
	// Auther is user from the foreign key
	createBlogTable = `CREATE TABLE IF NOT EXISTS article (
		post_id BIGINT PRIMARY KEY AUTO_INCREMENT,
		title VARCHAR(255) NOT NULL,
		content TEXT,
		author VARCHAR(255),
		Thumbnail BLOB,
		PublishOn DATETIME DEFAULT CURRENT_TIMESTAMP,
		ModifiedOns DATETIME DEFAULT CURRENT_TIMESTAMP,
		PRIMARY KEY (post_id),
		FOREIGN KEY (author) REFERENCES user(id)
	);`

	userProfile = `CREATE TABLE profile`
)
