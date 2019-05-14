package kunsul

import (
	"github.com/Masterminds/sprig"
	"k8s.io/api/extensions/v1beta1"
	"k8s.io/client-go/rest"
	"net/http"
	"html/template"

	log "github.com/Sirupsen/logrus"
)

type PageData struct {
	Title     string
	Ingresses []v1beta1.Ingress
}

func render(w http.ResponseWriter, r *http.Request, rest *rest.Config) {
	var ingresses []v1beta1.Ingress
	if ingresses, err = GetIngresses(rest); err != nil {
		returnError(w,err)
		return
	}
	log.Debugf("INGRESSES:>  %s", ingresses)

	var tmpl *template.Template

	//t := template.New("base").Funcs(sprig.FuncMap()).ParseFiles("template.html")
	templateName := "template.html"
	if tmpl, err = template.New(templateName).Funcs(sprig.FuncMap()).ParseFiles(templateName); err != nil {
		returnError(w,err)
		return
	}
	log.Debugf("TEMPLATES:>  %s", tmpl.DefinedTemplates())
	pageData := PageData{
		Title:     "kunsul",
		Ingresses: ingresses,
	}
	log.Debugf("PAGEDATA:>  %s",pageData)
	log.Debugf("PAGEDATA INGRESSES:> %s", pageData.Ingresses)

	if err := tmpl.Execute(w, pageData); err != nil {
		returnError(w,err)
		return
	}
}


func returnError(w http.ResponseWriter, e error ){
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(e.Error()))
	return
}