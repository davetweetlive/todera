package query

const (
	CreateUserTable = `CREATE TABLE IF NOT EXISTS user (
		user_id BIGINT PRIMARY KEY AUTO_INCREMENT,
		username VARCHAR(255) NOT NULL,
		email  VARCHAR(255) NOT NULL,
		password TEXT NOT NULL,
		first_name VARCHAR(255),
		last_name VARCHAR(255),
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		is_superuser BOOL, 
		profile_photo BLOB);`

	CreateArticleTable = `CREATE TABLE IF NOT EXISTS article (
		post_id BIGINT PRIMARY KEY AUTO_INCREMENT,
		title VARCHAR(255) NOT NULL,
		content TEXT,
		author BIGINT,
		thumbnail BLOB,
		release_time DATETIME DEFAULT CURRENT_TIMESTAMP,
		update_time DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (author) REFERENCES user(user_id));`

	CreateCommentTable = `CREATE TABLE IF NOT EXISTS comment (
		comment_id BIGINT PRIMARY KEY AUTO_INCREMENT,
		article BIGINT,
		comment TEXT ,
		comment_on DATETIME DEFAULT CURRENT_TIMESTAMP,
		update_time DATETIME DEFAULT CURRENT_TIMESTAMP,
		user BIGINT,
		FOREIGN KEY (user) REFERENCES user(user_id),
		FOREIGN KEY (article) REFERENCES article(post_id));`

	CreateTagTable = `CREATE TABLE IF NOT EXISTS tag (
		tag_id BIGINT PRIMARY KEY AUTO_INCREMENT,
		tag_name BIGINT);`
)

// Queries to create rows in different tables
// User table
// Article tables
// Comment table
// Tag Table
const (
	CreateUser = `INSERT INTO user (username, password, email, 
		is_superuser) VALUES (?, ?, ?, ?);`

	CreatePost = ``

	CreateComment = ``

	CreateTag
)

// Utilities queries
const (
	IfEmailExists    = `SELECT email FROM user WHERE email = ?;`
	IfUsernameExists = `SELECT username FROM user WHERE username = ?;`
)
