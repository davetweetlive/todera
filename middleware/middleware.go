package middleware

import (
	"madhyam/sessions"
	"net/http"
)

func AuthRequired(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := sessions.Store.Get(r, "session")

		_, ok := session.Values["username"]
		if !ok {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		handler.ServeHTTP(w, r)
	}
}
