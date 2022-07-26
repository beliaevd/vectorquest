package middleware

import (
	"Vectorquest/handler"
	"Vectorquest/tools/helper"
	"net/http"
	"time"
)

func CheckAccess(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("sessionId")
		if err != nil {
			http.Redirect(w, r, "/admin/login", 302)
			return
		}
		session, err := helper.GetAdminSession(cookie.Value)

		if err != nil {
			handler.AdminIndex(w, r)
			return
		}

		if session.LifeTime > time.Now().Unix() {
			handler.AdminIndex(w, r)
			return
		}

		if session.UserStatus == "N" {
			handler.AdminAccessDenied(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func CheckAdminUserAccess(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("sessionId")
		if err != nil {
			http.Redirect(w, r, "/admin/login", 302)
			return
		}
		session, err := helper.GetAdminSession(cookie.Value)

		if err != nil {
			handler.AdminIndex(w, r)
			return
		}
		if session.UserStatus == "N" {
			handler.AdminAccessDenied(w, r)
			return
		}
		if session.LifeTime > time.Now().Unix() {
			handler.AdminIndex(w, r)
			return
		}
		for i := 0; i < len(session.UserCat); i++ {
			if session.UserCat[i] == "Creater" {
				next.ServeHTTP(w, r)
				return
			}
		}
		handler.AdminAccessDenied(w, r)
		return
	})
}
