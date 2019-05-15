package kunsul

import (
	"net/http"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// we are always healthy!
	writeJsonResponse(w, http.StatusOK, `{"healthy": true}`)

	// let's not log health checks
	// logRequest(r)
}
