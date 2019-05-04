package kunsul

import (
	"fmt"
	"html/template"
	"k8s.io/api/extensions/v1beta1"
	"k8s.io/client-go/rest"
	"net/http"

	log "github.com/Sirupsen/logrus"
)

type PageData struct {
	Title     string
	Ingresses []v1beta1.Ingress
}

func render(w http.ResponseWriter, r *http.Request, rest *rest.Config) {
	ingresses := GetIngresses(rest)
	log.Debug(ingresses)

	tmpl, err := template.ParseFiles("template.html")
	if err != nil {
		panic(err)
	}
	log.Debug(tmpl)
	pageData := PageData{
		Title:     "kunsul",
		Ingresses: ingresses,
	}
	log.Debug(pageData)
	log.Debug(fmt.Sprintf("INSPECT %s", pageData.Ingresses))
	tmpl.Execute(w, pageData)
}
