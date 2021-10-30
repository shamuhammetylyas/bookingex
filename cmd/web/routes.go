package main

import (
	"net/http"

	"github.com/ShamuhammetYlyas/bookings/internal/config"
	"github.com/ShamuhammetYlyas/bookings/internal/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	// mux := pat.New()
	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	// main.go-da server ishledilende gelyan requestleri main package-in routes funksiyasy handle eder diyip belledik
	// bu yerde bolsa DefaultServeMux ulanman third party package ulandyk
	mux := chi.NewRouter()

	//mux.Use bilen route-lerde ulanyljak middleware-lerimizi bildiryaris.
	// her requestde shu middleware-ler ishleyar
	// middleware.Recoverer middleware-i package bilen gelyan middleware
	mux.Use(middleware.Recoverer)

	// NoSurf middleware-i middleware.go-da doreden oz middlewaremiz
	// Bu middleware CSRF ucin
	mux.Use(NoSurf)

	// SessionLoad hem oz doreden middlewaremiz.
	// biz bir requestde session doredenmiz bilen sho doredilen sessiony bashga request page tananok
	// meselem biz home page-da session doretsek, son about page-da sho session datany aljak bolsak alyp bilmeyaris
	// sebabi about page tananok. shony hemme request pageler tanar yaly etjek bolsak middleware ulanmaly bolyarys
	// SessionLoad diyip custom middleware doretdik, on icinde bolsa LoadAndSave diyip funksiya bar
	mux.Use(SessionLoad)

	//localhost:8080-e request ugratsa handlers package-in Repo structyn home handleri ishleyar
	//localhost:8080/about request ugratsa handlers package-in Repo structyn about handleri ishleyar
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/contact", handlers.Repo.Contact)
	mux.Get("/generals-quarters", handlers.Repo.Generals)
	mux.Get("/majors-suite", handlers.Repo.Majors)
	mux.Get("/make-reservation", handlers.Repo.Reservation)
	mux.Post("/make-reservation", handlers.Repo.PostReservation)
	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	mux.Post("/search-availability-json", handlers.Repo.AvailabilityJSON)

	// web sahypada goyulan static fayllary(img, css, js) sho web sahypalarda enable etmek ucin fileServer doretmeli
	// bu fileServer file-leri serve etmek ucin bir manager diyip hasap edelin. FileServer bir fileSystem-e garashyar
	// yagny haysy folderin icindaki fileleri serve etjekdigimizi bildiryaris. biz shu yerde root directory(bookings)
	// icindaki static folderin icinden file serve etjekdigimizi bildiryaris.
	fileServer := http.FileServer(http.Dir("./static/"))

	// bu yerdaki mux.Handle hem static bilen baslayan requestleri
	// localhost:8080/static bolup bashlayan requestleri handle etmek ucin ulanylyar
	// eger biz localhost:8080/static/GoWebProgramming.pdf diyip request ugratsak go sho fayly
	// (aslynda htmlin icinde <img src="static/GoWebProgramming.pdf" alt="image"> diyip bir zat bolsa sho hem request hasap edilyar.)
	// /static/static/GoWebProgramming.pdf directorynyn icinden gozleyar sebabi biz yokarda static folderin icinden
	// file serve etjekdigimizi bildirdik. Bu yagdayda bizin static folderimizin icinde static folder yok bolany ucin file tapylmayar
	// Shonun ucin file serve edilmaka static/static/GoWebProgramming.pdf -den biz ortadaky static-i ayyrmaly bolyarys.
	// Onun ucin http.StripPrefix() ulanylyar
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
