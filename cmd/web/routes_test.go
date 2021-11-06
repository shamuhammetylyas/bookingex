package main

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/ShamuhammetYlyas/bookings/internal/config"
)

// routes.go-daky routes funksiyasyny test etmek ucin ulanylyar
// routes funksiyasy config.AppConfig typeli pointere garashyany ucin shony doredip ugratdyk
func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)
	switch v := mux.(type) {
	case http.Handler:
		//do nothing
	default:
		t.Error(fmt.Sprintf("Type is not http.Handler, but is %T", v))
	}
}
