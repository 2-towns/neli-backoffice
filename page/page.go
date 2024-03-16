package page

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"gitlab.com/arnaud-web/neli-backoffice/config"
	"gitlab.com/arnaud-web/neli-backoffice/models"
)

func create(r *http.Request) *template.Template {
	var err error
	tpl, err := template.New("").Funcs(template.FuncMap{
		"api": func() string {
			return *config.ApiURL
		},
		"account": func() *models.Account {
			ctx := r.Context()
			a, _ := ctx.Value("account").(*models.Account)
			return a
		},
	}).ParseGlob("./static/templates/base/*.gohtml")

	if err != nil {
		log.Fatal(err)
	}

	tpl, err = tpl.ParseGlob(fmt.Sprintf("./static/templates/%s/*.gohtml", r.URL.Path))

	if err != nil {
		log.Fatal(err)
	}

	return tpl
}

// Display a template with data
func Display(w http.ResponseWriter, r *http.Request, name string, data interface{}) {
	tpl := create(r)
	err := tpl.ExecuteTemplate(w, name, data)

	if err != nil {
		log.Print(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
