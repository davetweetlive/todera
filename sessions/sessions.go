package sessions

import "github.com/gorilla/sessions"

var key = []byte("Saffron-Extention-powered-by-Saffron-Coders-Secret-Key")
var Store = sessions.NewCookieStore(key)
