package controllers

import (
	"crypto/subtle"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/admin") {
			if cookie, err := r.Cookie("auth"); err != nil {
				if err == http.ErrNoCookie {
					redirectToLogin(w, r)
					return
				} else {
					internalError(err, w)
					return
				}
			} else {
				if subtle.ConstantTimeCompare([]byte(cookie.Value), authArray) == 1 {
					next.ServeHTTP(w, r)
					return
				} else {
					redirectToLogin(w, r)
					return
				}
			}
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

func redirectToLogin(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login?next="+r.URL.Path, http.StatusFound)
}
