package routes

import (
	"fmt"
	"madhyam/models/query"
	"net/http"
)

func UserTable() {
	row, err := db.Query(query.CreateUser)
	if err != nil {
		fmt.Println("Can't run the create user query", err)
	}
	row.Close()
}

func ArticleTable() {
	row, err := db.Query(query.CreateArticle)
	if err != nil {
		fmt.Println("Can't run the create article query", err)
	}
	row.Close()
}

func CommentTable() {
	row, err := db.Query(query.CreateComment)
	if err != nil {
		fmt.Println("Can't run the create comment query", err)
	}
	row.Close()
}

func TagTable() {
	row, err := db.Query(query.CreateTag)
	if err != nil {
		fmt.Println("Can't run the create tag query", err)
	}
	row.Close()
}

func MigrationHandler(w http.ResponseWriter, r *http.Request) {
	UserTable()
	ArticleTable()
	CommentTable()
	TagTable()
	w.Write([]byte("Migrated succesfully"))
}
