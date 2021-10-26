package handlers

import (
	"net/http"

	"github.com/ShamuhammetYlyas/bookings/pkg/config"
	"github.com/ShamuhammetYlyas/bookings/pkg/models"
	"github.com/ShamuhammetYlyas/bookings/pkg/render"
)

//Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
//Repository pattern ulanmagymyzyn sebabi app configurationlaryn hem handlers package-de
// hemde render package-de ulanylyandygy ucin. Kop yerde app config gerek bolany ucin
// ony bir reponyn icine salmak maslahat berilyar
type Repository struct {
	App *config.AppConfig
}

//NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(res http.ResponseWriter, req *http.Request) {
	remoteIP := req.RemoteAddr
	m.App.Session.Put(req.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(res, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(res http.ResponseWriter, req *http.Request) {
	//some logic here
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again"

	remoteIP := m.App.Session.GetString(req.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(res, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
