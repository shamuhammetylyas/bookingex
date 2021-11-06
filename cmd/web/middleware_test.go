package main

import (
	"fmt"
	"net/http"
	"testing"
)

// middleware.go-daky NoSurf funksiyany test etmek ucin doredildi
// middleware.go-daky NoSurf funksiyasy parametr hokmunda http.Handler interface-in memberi bolan bir type-li value
// garashyany ucin biz test etjek bolamyzda hem sho requirementlerini yerine getirmeli.
// onun ucin bolsa birinji http.Handler interface-in memberini doretmeli.
// test ishlemaka bir zatlar setup etjek bolsak shony setup_test.go funksiyanyn icinde edyaris
// shon icindaki kodlar her test ishlap bashlamaka birinji sholar ishleyar
func TestNoSurf(t *testing.T) {
	//setup_test.go-da bir custom type doretdik we ona ServeHTTP receiver funksiyasyny berdik
	//indi ashakdaky myH http.Handler interface-in bir memberi bolany ucin biz ony NoSurf metodyna iberip bilyaris.
	var myH myHandler
	h := NoSurf(&myH)

	// middleware.go-daky NoSurf bize http.Handler return edyar
	// eger http.Handler return edilman bashga bir zat return edilse onda test fail boldugy bolyar
	// shony kontrol etmek ucin hem ashakdaky kod ulanylsa bolyar
	switch v := h.(type) {
	case http.Handler:
		//do nothing
	default:
		t.Error(fmt.Sprintf("Type is not http.Handler, but is %T", v))
	}
}

// middleware.go-daky SessionLoad funksiyasyny test etmek ucin ulanylyar

func TestSessionLoad(t *testing.T) {
	//setup_test.go-da bir custom type doretdik we ona ServeHTTP receiver funksiyasyny berdik
	//indi ashakdaky myH http.Handler interface-in bir memberi bolany ucin biz ony SessionLoad metodyna iberip bilyaris.
	var myH myHandler
	h := SessionLoad(&myH)

	// middleware.go-daky SessionLoad bize http.Handler return edyar
	// eger http.Handler return edilman bashga bir zat return edilse onda test fail boldugy bolyar
	// shony kontrol etmek ucin hem ashakdaky kod ulanylsa bolyar
	switch v := h.(type) {
	case http.Handler:
		//do nothing
	default:
		t.Error(fmt.Sprintf("Type is not http.Handler, but is %T", v))
	}
}
