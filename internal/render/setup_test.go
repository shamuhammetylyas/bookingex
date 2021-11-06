package render

import (
	"encoding/gob"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/ShamuhammetYlyas/bookings/internal/config"
	"github.com/ShamuhammetYlyas/bookings/internal/models"
	"github.com/alexedwards/scs/v2"
)

var session *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M) {
	gob.Register(models.Reservation{})

	testApp.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	testApp.Session = session

	// ashakdaky app render packagein icindaki app variable-dir
	// bir AddDefaultData-da app ulanyanymzy ucin test bashlamak bir app configurationyny
	// doredyaris, we render.go-daky app vairable-a beryaris
	// son meselem render.go-daky AddDefaultData-da app ulanylyar.
	// sho yerde yazylan app ashakdaky app-dir
	// aslynda sho yerde ulanylyan app pointer testApp-in pointeridir.
	// bu dine TEST-de
	app = &testApp
	os.Exit(m.Run())
}

type myWriter struct{}

func (tw *myWriter) Header() http.Header {
	var h http.Header
	return h
}

func (tw *myWriter) WriteHeader(i int) {}

func (tw *myWriter) Write(b []byte) (int, error) {
	lenght := len(b)
	return lenght, nil
}
