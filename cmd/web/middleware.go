package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// NoSurf adds CSRF protection to all POST requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// LoadAndSave loads and saves the session on every request.
// Doredilen sessiony load edip shony her request-e save edyar.
// Muny etmek bilen biz her sahypada(handlerde) doredilen session valueny alyp bilyaris
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
