package kunsul

import (
	"io"
	"net/http"
)

func writeJsonResponse(w http.ResponseWriter, status int, body string) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, body)
}

func writeHtmlErrorResponse(w http.ResponseWriter, e error ){
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(e.Error()))
	return
}