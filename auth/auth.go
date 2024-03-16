package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"gitlab.com/arnaud-web/neli-backoffice/api"
	"gitlab.com/arnaud-web/neli-backoffice/config"
	"gitlab.com/arnaud-web/neli-backoffice/models"
	"gitlab.com/arnaud-web/neli-backoffice/page"
)

const (
	SuperAdminRole = "super_admin"
	adminRole      = "administrator"
)

func Login(w http.ResponseWriter, r *http.Request) {
	_, err := r.Cookie("access_token")

	if err == nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	var i interface{}
	page.Display(w, r, "login.gohtml", i)
}

func StoreInfo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("access_token")

		// If cookie doesn't exist move to next request
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		a := new(models.Account)
		url := fmt.Sprintf("%s/users/me", *config.ApiURL)

		if err := api.Get(url, c.Value, a); err != nil {
			log.Println(err)
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "account", a)
		next.ServeHTTP(w, r.WithContext(ctx))

	})
}

func checkRole(next http.Handler, role string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		a, ok := ctx.Value("account").(*models.Account)

		if !ok {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}

		if !strings.Contains(role, a.Role) {
			http.Error(w, "Bad role", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
func SuperAdmin(next http.Handler) http.Handler {
	return checkRole(next, SuperAdminRole)
}

func Admin(next http.Handler) http.Handler {
	return checkRole(next, adminRole)
}

func SuperAdminOrAdmin(next http.Handler) http.Handler {
	return checkRole(next, fmt.Sprintf("%s,%s", SuperAdminRole, adminRole))
}
