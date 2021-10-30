package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/ShamuhammetYlyas/bookings/internal/config"
	"github.com/ShamuhammetYlyas/bookings/internal/forms"
	"github.com/ShamuhammetYlyas/bookings/internal/models"
	"github.com/ShamuhammetYlyas/bookings/internal/render"
)

//Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
// Repository pattern ulanmagymyzyn sebabi app configurationlaryn hem handlers package-de
// hemde render package-de ulanylyandygy ucin. Kop yerde app config gerek bolany ucin
// biz handler package ucin repository pattern ulandyk. Hokman shuny ulanmaly diyen zat yok
// App-i main.go-dan NewRepo funksiyasyna gelyan app-in adresine denledik
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
// main.go-da shu NewRepo, doredilen app configin adresini ugradypdyk.
// bu funksiya hem bir repository doredip shonun adresini return edyar.
// main.go-da hem shu doredilen repositoryn adresini alyp bir repo variable-a
// denledik yagny main.go-da repo=0xc213123123(doredilen repositoryn adresi)
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//NewHandlers sets the repository for the handlers
//NewHandlers funskiyasy main.go-da ulanylyar. Bu metod NewRepodan return edilen repositoryn adresine garashyar
//we sho adresi alyp yokarda doredilen Repo variable-a denleyar. yagny shuwagt Repo = 0xca1231239(doredilen reponyn adresi)
func NewHandlers(r *Repository) {
	Repo = r
}

//Home handleri indi repositoryn receiver funksiya boldy
// biz nirede Home handleri ulanjak bolsak Repo objectin usti bilen ulanmaly bolyarys.
func (m *Repository) Home(res http.ResponseWriter, req *http.Request) {
	// req.RemoteAddr request ugradyan clientin adresini beryar.
	// ony alyp remoteIP variable-a denleyaris
	remoteIP := req.RemoteAddr

	// m, doredilen Repo.
	// App sho reponyn App propertisi.
	// Bu propertinin type-i hem *config.AppConfig bolany ucin onun icindaki propertyleri ulanyp bolyar property
	// yagny structyn icinde struct typeli property bar
	// Session *config.AppConfigin Session propertisi. Biz muny main.go-da beripdik ilki bashda
	// Put session managerin receiver funksiyasy
	m.App.Session.Put(req.Context(), "remote_ip", remoteIP)
	render.RenderTemplate(res, req, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(res http.ResponseWriter, req *http.Request) {
	//some logic here
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello again"

	remoteIP := m.App.Session.GetString(req.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(res, req, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) Contact(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, req, "contact.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Generals(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, req, "generals.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Majors(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, req, "majors.page.tmpl", &models.TemplateData{})
}

func (m *Repository) Reservation(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, req, "make-reservation.page.tmpl", &models.TemplateData{
		// laravelidaki formdaky old value-lar ucin we formyn errorlaryny gorkezmek ucin
		// renderde-de form objecti gerek bolyar
		Form: forms.New(nil),
	})
}

func (m *Repository) PostReservation(res http.ResponseWriter, req *http.Request) {

}

func (m *Repository) Availability(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, req, "search-availability.page.tmpl", &models.TemplateData{})
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func (m *Repository) AvailabilityJSON(res http.ResponseWriter, req *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Available",
	}

	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Println(err)
	}

	res.Header().Set("Content-Type", "application/json")
	res.Write(out)

}

func (m *Repository) PostAvailability(res http.ResponseWriter, req *http.Request) {

}
