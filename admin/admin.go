package admin

import (
	"net/http"

	"gitlab.com/arnaud-web/neli-backoffice/page"
)

func Get(w http.ResponseWriter, r *http.Request) {
	var i interface{}
	page.Display(w, r, "admin.gohtml", i)
}

func Password(w http.ResponseWriter, r *http.Request) {
	var i interface{}
	page.Display(w, r, "password.gohtml", i)
}

func Info(w http.ResponseWriter, r *http.Request) {
	var i interface{}
	page.Display(w, r, "info.gohtml", i)
}

func Reset(w http.ResponseWriter, r *http.Request) {
	var i interface{}
	page.Display(w, r, "reset.gohtml", i)
}

func Content(w http.ResponseWriter, r *http.Request) {
	var i interface{}
	page.Display(w, r, "content.gohtml", i)
}

func Leader(w http.ResponseWriter, r *http.Request) {
	var i interface{}
	page.Display(w, r, "leader.gohtml", i)
}

func MaxDuration(w http.ResponseWriter, r *http.Request) {
	var i interface{}
	page.Display(w, r, "max-duration.gohtml", i)
}

func Cleaning(w http.ResponseWriter, r *http.Request) {
	var i interface{}
	page.Display(w, r, "cleaning.gohtml", i)
}
