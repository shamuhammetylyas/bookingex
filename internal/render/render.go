package render

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/ShamuhammetYlyas/bookings/internal/config"
	"github.com/ShamuhammetYlyas/bookings/internal/models"
	"github.com/justinas/nosurf"
)

var functions = template.FuncMap{}
var app *config.AppConfig

// shu ashakdaky ./templates render.go-da bolany ucin bookings yagny root folderin icindaki templates folder diyildigi bolyar
// yone test-de bu uytgemeli
var pathToTemplates = "./templates"

// AddDefaultData-daky datalar hemme template-a gidyar. bosh yada doly
func AddDefaultData(td *models.TemplateData, req *http.Request) *models.TemplateData {
	// usere message ibermek ucin gowy.
	// PopString sessionda duran valueny alyarda(key-ine gora) sessiondan pozyar
	td.Flash = app.Session.PopString(req.Context(), "flash")
	td.Error = app.Session.PopString(req.Context(), "error")
	td.Warning = app.Session.PopString(req.Context(), "warning")
	// CSRFToken datasy default data. Hemme template-e shu data gidyar
	td.CSRFToken = nosurf.Token(req)
	return td
}

// NewRenderer sets the app config for template package
// app configlerini render package-da ulanmak ucin
// bu yerde-de doredilen app configin adresini alyp app variable-a denledik
// app=0xc123453453(main.go-da doredilen app config adresi)
func NewRenderer(a *config.AppConfig) {
	app = a
}

func Template(res http.ResponseWriter, req *http.Request, tmpl string, td *models.TemplateData) error {
	// eger app production mod-da bolsa onda app configin icindaki parse edilen templateler ulanylyar
	// app configin icinde
	// app compile edilende templateler parse edilyarde app configin icindaki TemplateCache-de saklanyar
	// programma shu yere gelmaka app.TemplateCache = {'home.page.tmpl' : 0xada123143(parse edilen home templatin adresi)}
	// sonra render-den home.page.tmpl gelende shu yer app-in useCache-ne seredyar. eger true bolsa onda parse edilen template-i
	// app configin TemplateCache-inden alyar. eger false bolsa templateler tazeden parse edilyar. Development mod-da tazedilen parse
	// edileni gerek bize. Parse edyarde tc variable-a beryar
	var tc map[string]*template.Template
	if app.UseCache {
		//get the template cache from app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	//shu wagt tc= {'home.page.tmpl' : 0xada123143(parse edilen home templatin adresi), 'avout.page.tmpl' : 0xada123asd143(parse edilen about templatin adresi)}
	// tmpl render-den gelyan template-in ady. home.page.tmpl gelse shony alyarda t variable-a beryar
	t, ok := tc[tmpl]
	if !ok {
		// log.Fatal("Could not get template from template cache")
		// log.Println("can't get template from cache")
		return errors.New("can't get template from cache")
	}

	// td renderden gelyan data.
	// eger biz bir datany hemme template-da ulanjak bolsak template execute edilmaka shona sho datany goshmaly bolyarys.
	// shu yerde renderden gelyan data yenede ussune data goshyarys.
	td = AddDefaultData(td, req)

	// bytes package-in icindaki Buffer structy ucin ramda bir yer allocate edyar
	buf := new(bytes.Buffer)

	// parse ediljek template-i res-e execute etman buf-e execute edyar.
	// sebabi biz goni t.Execute(res, td) etsek browserde hic zat gorunmeyar.
	// cunki biz t variable-a template pointeri diskden parse edip almadyk
	// yagny template.parseFiles funksiyasy bermedi sho pointeri. tc-nin ussi bilen aldyk
	// shular yaly yagdayda biz biraz bashgacaarak execute etmeli bolyarys.
	// template-i buf-a yazdyryarys. sonra bolsa bufun WriteTo receiver funksiyasy bilen res-e yazdyryarys.
	_ = t.Execute(buf, td)
	_, err := buf.WriteTo(res)
	if err != nil {
		// log.Fatal(err)
		fmt.Println("Error writing template to browser", err)
		return err
	}

	return nil
}

// CreateTemplateCache creates a template cache as a map
// proje compile edilende, server bashlamaka template-leri parse edip bir yerde saklamaly bolyar
// shonun ucin main.go-da server bashlamaka onundaki kodlar ishar yaly onunden yazylyar.
// web server bashlamaka template-ler parse edilyar shu yerde

func CreateTemplateCache() (map[string]*template.Template, error) {

	// parse edilen template-ler bir yerde key-value gornushde saklanmaly
	// shonun ucin shu map doredildi.
	myCache := map[string]*template.Template{}

	// templates folderin icinde adynda .page.tmpl bar bolan fayllary path-i bilen alyar
	// string slice gaytaryar.
	// ["templates/about.page.tmpl", "templates/home.page.tmpl"]
	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplates))
	if err != nil {
		return myCache, err
	}

	// yokardaky slice-i for loop bilen loop edyar
	for _, page := range pages {

		// page shu yerde "templates/about.page.tmpl"
		// bize bolsa dine about.page.tmpl gerek.
		// shonun ucun Base funksiyasy ulanylyar. dine adyny alyar we ony name-a assign edyar
		// name = about.page.tmpl (1-nji loopda)
		name := filepath.Base(page)

		// template.New("about.page.tmpl") adynda bir template doredyar we onun
		// icinde kabir funksiyalar ulanyljakdygyny aydyar. we sho doredilen template-a hem
		// templates/about.page.tmpl-i parse edip icine yerleshdiryar
		// yagny about.page.tmpl adyndaky template-de templates/about.page.tmpl-in parse edileninden sonky adresi duryar
		// ony hem ts-e beryar
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		// templates folderin icinde adynda .layout.tmpl bar bolan fayllary path-i bilen alyar
		// string slice gaytaryar.
		// ["templates/base.layout.tmpl"]
		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
		if err != nil {
			return myCache, err
		}

		// eger .layout.tmpl-li bir fayl bar bolsa onda
		// yanky ts-in icindaki duran template-in ussune layouty hem parse edyar
		// netije-de layout bilen page merge bolan bolyar.
		// biz yone ts.Execute edenimiz bilen layout hem parse boldygy bolanok.
		// shonun ucin parse edilen template-in ussune layouty hem parse edip birleshdirmeli
		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
			if err != nil {
				return myCache, err
			}
		}

		// loop 2 gezek gaytalanyar netijede
		// myCache["about.page.tmpl"] = 0xcad21231jkj13(layout bilen birleshdirilen about page)
		// myCache["home.page.tmpl"] = 0xcad21231jkj13(layout bilen birleshdirilen home page)
		myCache[name] = ts
	}

	return myCache, nil
}
