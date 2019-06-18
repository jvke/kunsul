package kunsul

import (
	"path/filepath"
	"html/template"
	"net/http"

	"github.com/masterminds/sprig"
	"k8s.io/api/core/v1"
	"k8s.io/api/extensions/v1beta1"
	"k8s.io/client-go/rest"
	log "github.com/sirupsen/logrus"
)

type PageData struct {
	Title     string
	Ingresses []v1beta1.Ingress
	Services  []v1.Service
}

var (
	tmpl *template.Template
)

func render(w http.ResponseWriter, r *http.Request, rest *rest.Config, configDir string, templateFile string) {
	var ingresses []v1beta1.Ingress
	var services []v1.Service

	if ingresses, err = GetIngresses(rest); err != nil {
		writeHtmlErrorResponse(w, err)
		return
	}
	log.Debugf("INGRESSES:>  %s", ingresses)

	if services, err = GetServices(rest); err != nil {
		writeHtmlErrorResponse(w, err)
		return
	}
	log.Debugf("SERVICES:>  %s", services)

	if tmpl, err = template.New(filepath.Base(templateFile)).Funcs(sprig.FuncMap()).ParseFiles(templateFile); err != nil {
		writeHtmlErrorResponse(w, err)
		return
	}
	log.Debug(tmpl.DefinedTemplates())

	pageData := PageData{
		Ingresses: ingresses,
		Services:  services,
	}

	log.Debugf("PAGEDATA:>  %s", pageData)
	log.Debugf("PAGEDATA INGRESSES:> %s", pageData.Ingresses)

	if err := tmpl.Execute(w, pageData); err != nil {
		writeHtmlErrorResponse(w, err)
		log.Debugf("errored template: %s", tmpl)
		return
	}
}
