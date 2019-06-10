package kunsul

import (
	"fmt"
	"k8s.io/client-go/rest"
	"net/http"

	log "github.com/sirupsen/logrus"
)

var (
	code int = http.StatusTemporaryRedirect
)

func Serve(kubeConfig *rest.Config, configDir string, template string, listenPort int, listings bool, accessLog bool) {
	log.Info("initialize kunsul")
	log.Debug("debug logging enabled")

	http.HandleFunc("/health", healthCheckHandler)
	http.HandleFunc("/health/", healthCheckHandler)

	log.WithFields(log.Fields{
		"configDir":  configDir,
		"port": listenPort,
	}).Info("listening for requests")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		render(w, r, kubeConfig, configDir, template)
	})

	http.ListenAndServe(":"+fmt.Sprintf("%v", listenPort), nil)
}
