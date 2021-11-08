package helpers

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/ShamuhammetYlyas/bookings/internal/config"
)

var app *config.AppConfig

func NewHelpers(a *config.AppConfig) {
	app = a
}

func ClientError(res http.ResponseWriter, status int) {
	app.InfoLog.Println("Client error with status of ", status)
	http.Error(res, http.StatusText(status), status)
}

func ServerError(res http.ResponseWriter, err error) {
	// err.Error errory string edip yazyar
	// debug.Stack bolsa error barada biraz kopurak maglumat beryar
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	http.Error(res, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
