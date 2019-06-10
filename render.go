package kunsul

import (
	"github.com/Masterminds/sprig"
	log "github.com/sirupsen/logrus"
	"html/template"
	"k8s.io/api/extensions/v1beta1"
	"k8s.io/client-go/rest"
	"net/http"
	"path"
)

type PageData struct {
	Title     string
	Ingresses []v1beta1.Ingress
}

func render(w http.ResponseWriter, r *http.Request, rest *rest.Config, configDir string, templateName string) {
	var ingresses []v1beta1.Ingress
	if ingresses, err = GetIngresses(rest); err != nil {
		writeHtmlErrorResponse(w,err)
		return
	}
	log.Debugf("INGRESSES:>  %s", ingresses)

	var tmpl *template.Template

	//t := template.New("base").Funcs(sprig.FuncMap()).ParseFiles("template.html")
	var tmplPath = path.Join(configDir,templateName)
	if tmpl, err = template.New(templateName).Funcs(sprig.FuncMap()).ParseFiles(tmplPath); err != nil {
		writeHtmlErrorResponse(w,err)
		return
	}
	log.Debugf("TEMPLATES:>  %s", tmpl.DefinedTemplates())
	pageData := PageData{
		Ingresses: ingresses,
	}
	log.Debugf("PAGEDATA:>  %s",pageData)
	log.Debugf("PAGEDATA INGRESSES:> %s", pageData.Ingresses)

	if err := tmpl.Execute(w, pageData); err != nil {
		writeHtmlErrorResponse(w,err)
		return
	}
}


