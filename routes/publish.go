package routes

import (
	"fmt"
	"io/ioutil"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

func PublishPostBlog(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// Upload a file
	file, fh, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	fmt.Println("File handler", fh.Filename)

	text, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}
	w.Write(text)
}

func WriteCookie(w http.ResponseWriter, r *http.Request) {
	// http.SetCookie(w, &http.Cookie{
	// 	Name:  "my-cookie",
	// 	Value: "some value",
	// })
	// fmt.Fprintln(w, "COOKIE WRITTNE - CHECK YOUR BROWSER")
	// fmt.Fprintln(w, "CHECK APPlication in dev tools in chrome")
	cookie, err := r.Cookie("session")
	if err != nil {
		id, _ := uuid.NewV4()
		cookie := &http.Cookie{
			Name:  "session",
			Value: id.String(),
			// Secure: true,
			HttpOnly: true,
		}

		http.SetCookie(w, cookie)
	}

	fmt.Fprintln(w, cookie)
}

func ReadCookie(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("my-cookie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}
	fmt.Fprintln(w, "COOKIE WRITTNE ", c)
	// fmt.Fprintln(w, "CHECK APPlication in dev tools in chrome")
}
