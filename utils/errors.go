package utils

import (
	"madhyam/logging"
	"net/http"
)

func InternalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("Internal server error!"))
	logging.WriteLog(logging.FATAL, "Internal server error.")
}
