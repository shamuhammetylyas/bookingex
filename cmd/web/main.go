package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ShamuhammetYlyas/bookings/internal/config"
	"github.com/ShamuhammetYlyas/bookings/internal/driver"
	"github.com/ShamuhammetYlyas/bookings/internal/handlers"
	"github.com/ShamuhammetYlyas/bookings/internal/helpers"
	"github.com/ShamuhammetYlyas/bookings/internal/models"
	"github.com/ShamuhammetYlyas/bookings/internal/render"
	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8081"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	// shu yerde db driver package-in DB structynyn memberi
	// onun hem SQL propertisi bolany ucin, o property hem *sql.DB type-da bolany ucin
	// onun CLose funksiyasyny ulanyp bilyaris
	defer db.SQL.Close()
	defer close(app.MailChan)
	fmt.Println("Starting mail listener...")
	listenForMail()

	// msg := models.MailData{
	// 	To:      "john@do.ca",
	// 	From:    "me@here.com",
	// 	Subject: "Some subject",
	// 	Content: "",
	// }

	// app.MailChan <- msg

	fmt.Printf("Server started on port %s\n", portNumber)

	// serverimiz goshmaca configurationlar bilen ishleder yaly http.Server bilen ishledyaris
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

func run() (*driver.DB, error) {
	// what am I going to put in the session
	// sessionda data saklajakdygymyzy birinji applicationa bildirmeli
	// gob => built in package
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})

	mailChan := make(chan models.MailData)
	app.MailChan = mailChan

	//change this to true when in production
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "Error\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	// session manager doredilyar. Session manager sessionin ozi dal-de session doredyan manager
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	//sessiony XSS atakalardan goramak ucin ulanylyar. Development mod-da false duranynyn zyyany yok
	session.Cookie.Secure = app.InProduction

	// app configin seesionyna session manager berdik
	app.Session = session

	//connect to database
	log.Println("Connecting to database...")
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=booking user=postgres password=")
	if err != nil {
		log.Fatal("Cannot connect to database! Dying...")
	}

	log.Println("Connected to database")

	// hemme parse edilen template-leri tc variable-a assign etdik
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// parse edilen template-leri app configdaki TemplateCache-e beryaris
	// app.UseCache-i development mod-da false edenimizin sebabi render package-da template parse edilende
	// her gezek tazeden parse edilmegini isleyaris. muny render package-da ulanyarys.
	app.TemplateCache = tc
	app.UseCache = false

	// app configurationlaryny render package-da ulanmak ucin doreden app configimizn adresini render
	// package-in NewTemplate funksiyasyna iberyaris. Bu funksiya bolsa *config.AppConfig garashyar we gelen adresi bir
	// variable-a denleyar we netijede shu yerde doredilen app configurationymyz render packageda ulanar yaly bolar
	render.NewRenderer(&app)

	// app configurationlaryny handlers package-da ulanmak ucin doreden app configimizn adresini handler
	// package-in NewRepo funksiyasyna iberyaris. Bu funksiya bolsa *config.AppConfig garashyar we gelen adresi bir
	// variable-a denleyar we netijede shu yerde doredilen app configurationymyz handlers packageda ulanar yaly bolar

	// yokarda database connection acyldy, indi sho connectiony bashga packageler-de ulanar yaly bir zat etmeli
	// handlers package-da ulanmak ucin biz handler.go-da bir zatlar etdik. yone shonun ucin hem acylan db *driver.DB type gerek
	// shonun ucin ashakda NewRepo funksiyasyna parametr hokmunde ugradyarys
	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	helpers.NewHelpers(&app)
	return db, nil
}
