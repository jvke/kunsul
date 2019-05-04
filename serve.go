package kunsul

import (
	"fmt"
	"k8s.io/client-go/rest"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

var (
	code int = http.StatusTemporaryRedirect
)

func Serve(config *rest.Config, directory string, listenPort int, listings bool, accessLog bool) {
	log.Info("initialize kunsul")
	log.Debug("debug logging enabled")

	http.HandleFunc("/health", healthCheckHandler)
	http.HandleFunc("/health/", healthCheckHandler)

	log.WithFields(log.Fields{
		"dir":  directory,
		"port": listenPort,
	}).Info("listening for requests")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w, r, config)
	})

	http.ListenAndServe(":"+fmt.Sprintf("%v", listenPort), nil)
}
