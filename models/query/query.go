package query

const (
	CreateUser = `CREATE TABLE IF NOT EXISTS user (
		user_id BIGINT PRIMARY KEY AUTO_INCREMENT,
		username VARCHAR(255) NOT NULL,
		email  VARCHAR(255) NOT NULL,
		password TEXT NOT NULL,
		first_name VARCHAR(255),
		last_name VARCHAR(255),
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		is_superuser BOOL, 
		profile_photo BLOB);`

	CreateArticle = `CREATE TABLE article (
		post_id BIGINT PRIMARY KEY AUTO_INCREMENT,
		title VARCHAR(255) NOT NULL,
		content TEXT,
		author BIGINT,
		thumbnail BLOB,
		release_time DATETIME DEFAULT CURRENT_TIMESTAMP,
		update_time DATETIME DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (author) REFERENCES user(user_id));`

	CreateComment = `CREATE TABLE IF NOT EXISTS comment (
		comment_id BIGINT PRIMARY KEY AUTO_INCREMENT,
		article BIGINT,
		comment TEXT ,
		comment_on DATETIME DEFAULT CURRENT_TIMESTAMP,
		update_time DATETIME DEFAULT CURRENT_TIMESTAMP,
		user BIGINT,
		FOREIGN KEY (user) REFERENCES user(user_id),
		FOREIGN KEY (article) REFERENCES article(post_id));`

	CreateTag = `CREATE TABLE IF NOT EXISTS tag (
		tag_id BIGINT PRIMARY KEY AUTO_INCREMENT,
		tag_name BIGINT);`
)
