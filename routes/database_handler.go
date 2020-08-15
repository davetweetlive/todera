package routes

import (
	"fmt"
	"madhyam/models/query"
	"net/http"
)

func UserTable() {
	row, err := db.Query(query.CreateUserTable)
	if err != nil {
		fmt.Println("Can't run the create user query", err)
	}
	row.Close()
}

func ArticleTable() {
	row, err := db.Query(query.CreateArticleTable)
	if err != nil {
		fmt.Println("Can't run the create article query", err)
	}
	row.Close()
}

func CommentTable() {
	row, err := db.Query(query.CreateCommentTable)
	if err != nil {
		fmt.Println("Can't run the create comment query", err)
	}
	row.Close()
}

func TagTable() {
	row, err := db.Query(query.CreateTagTable)
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
