package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ShamuhammetYlyas/bookings/internal/config"
	"github.com/ShamuhammetYlyas/bookings/internal/handlers"
	"github.com/ShamuhammetYlyas/bookings/internal/models"
	"github.com/ShamuhammetYlyas/bookings/internal/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	// what am I going to put in the session
	// sessionda data saklajakdygymyzy birinji applicationa bildirmeli
	// gob => built in package
	gob.Register(models.Reservation{})

	//change this to true when in production
	app.InProduction = false

	// session manager doredilyar. Session manager sessionin ozi dal-de session doredyan manager
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	//sessiony XSS atakalardan goramak ucin ulanylyar. Development mod-da false duranynyn zyyany yok
	session.Cookie.Secure = app.InProduction

	// app configin seesionyna session manager berdik
	app.Session = session

	// hemme parse edilen template-leri tc variable-a assign etdik
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	// parse edilen template-leri app configdaki TemplateCache-e beryaris
	// app.UseCache-i development mod-da false edenimizin sebabi render package-da template parse edilende
	// her gezek tazeden parse edilmegini isleyaris. muny render package-da ulanyarys.
	app.TemplateCache = tc
	app.UseCache = false

	// app configurationlaryny render package-da ulanmak ucin doreden app configimizn adresini render
	// package-in NewTemplate funksiyasyna iberyaris. Bu funksiya bolsa *config.AppConfig garashyar we gelen adresi bir
	// variable-a denleyar we netijede shu yerde doredilen app configurationymyz render packageda ulanar yaly bolar
	render.NewTemplate(&app)

	// app configurationlaryny handlers package-da ulanmak ucin doreden app configimizn adresini handler
	// package-in NewRepa funksiyasyna iberyaris. Bu funksiya bolsa *config.AppConfig garashyar we gelen adresi bir
	// variable-a denleyar we netijede shu yerde doredilen app configurationymyz handlers packageda ulanar yaly bolar
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	fmt.Printf("Server started on port %s\n", portNumber)

	//serverimiiz goshmaca configurationlar bilen ishleder yaly http.Server bilen ishledyaris
	// http.Server bir structyr bir shu yerde structyn bir instance-ni doredip shon adresini srv variable-a beryaris
	// Handler bilen gelyan requestleri routes funksiyasy bilen handle etjekdigimizi bildiryaris.
	// onun icinde birden app configleri geerek bolan yagdayynda ulanar yaly doredilen app configin adresini ugdatyarys
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
