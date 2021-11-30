package controllers

import (
	"crypto/subtle"
	"github.com/m-butterfield/mattbutterfield.com/app/lib"
	"github.com/m-butterfield/mattbutterfield.com/app/static"
	"html/template"
	"net/http"
	"time"
)

var loginTemplatePath = append([]string{templatePath + "login.gohtml"}, baseTemplatePaths...)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		auth := r.FormValue("auth")
		if subtle.ConstantTimeCompare([]byte(auth), authArray) == 1 {
			http.SetCookie(w, &http.Cookie{Name: "auth", Value: auth})
			next := r.URL.Query().Get("next")
			if next != "" {
				http.Redirect(w, r, next, 302)
			} else {
				renderLoginPage(w, true)
			}
		} else {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
		}
	} else {
		renderLoginPage(w, isUserLoggedIn(r))
	}
}

func isUserLoggedIn(r *http.Request) bool {
	if cookie, err := r.Cookie("auth"); err != nil {
		return false
	} else if subtle.ConstantTimeCompare([]byte(cookie.Value), authArray) == 1 {
		return true
	}
	return false
}

func renderLoginPage(w http.ResponseWriter, loggedIn bool) {
	tmpl, err := template.ParseFS(&static.FlexFS{}, loginTemplatePath...)
	if err != nil {
		lib.InternalError(err, w)
		return
	}
	if err = tmpl.Execute(w, struct {
		Year     string
		LoggedIn bool
	}{
		Year:     time.Now().Format("2006"),
		LoggedIn: loggedIn,
	}); err != nil {
		lib.InternalError(err, w)
		return
	}
}
