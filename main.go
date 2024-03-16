package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi"
	"gitlab.com/arnaud-web/neli-backoffice/admin"
	"gitlab.com/arnaud-web/neli-backoffice/auth"
	"gitlab.com/arnaud-web/neli-backoffice/config"
	"gitlab.com/arnaud-web/neli-backoffice/models"
)

var logger = log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

func main() {
	logger.Println("Running server on port", *config.Port)

	r := chi.NewRouter()

	r.Use(auth.StoreInfo)

	r.With(auth.SuperAdminOrAdmin).Get("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		a, ok := ctx.Value("account").(*models.Account)

		if !ok {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		if a.Role == auth.SuperAdminRole {
			http.Redirect(w, r, "/admin", http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, "/leader", http.StatusSeeOther)
	})

	r.Get("/login", auth.Login)

	r.With(auth.SuperAdmin).Get("/admin", admin.Get)
	r.With(auth.SuperAdmin).Get("/max-duration", admin.MaxDuration)
	r.With(auth.SuperAdmin).Get("/cleaning", admin.Cleaning)

	r.With(auth.Admin).Get("/password", admin.Password)
	r.With(auth.Admin).Get("/info", admin.Info)
	r.With(auth.Admin).Get("/reset", admin.Reset)
	r.With(auth.Admin).Get("/content", admin.Content)
	r.With(auth.Admin).Get("/leader", admin.Leader)

	workDir, _ := os.Getwd()
	filesDir := filepath.Join(workDir, "static/assets")
	fileServer(r, "/static", http.Dir(filesDir))

	fmt.Println(http.ListenAndServe(fmt.Sprintf(":%d", *config.Port), r))
}

func fileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit URL parameters.")
	}

	fs := http.StripPrefix(path, http.FileServer(root))

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", 301).ServeHTTP)
		path += "/"
	}

	path += "*"

	r.Get(path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	}))

}
