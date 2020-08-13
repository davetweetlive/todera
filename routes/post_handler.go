package routes

import (
	"fmt"
	"net/http"
)

func TestTemplate(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "index.html", nil); err != nil {
		fmt.Println("Couldn't parse template")
	}
}

func CategoriesHandler(w http.ResponseWriter, r *http.Request) {
	if err := templates.ExecuteTemplate(w, "category.html", pageInfo); err != nil {
		fmt.Println("Couldn't parse category template")
	}
}
