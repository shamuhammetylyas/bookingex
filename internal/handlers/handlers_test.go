package handlers

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type postData struct {
	key, value string
}

var theTests = []struct {
	name               string
	url                string
	method             string
	params             []postData
	expectedStatusCode int
}{
	{"home", "/", "GET", []postData{}, http.StatusOK},
	{"about", "/about", "GET", []postData{}, http.StatusOK},
	{"gq", "/generals-quarters", "GET", []postData{}, http.StatusOK},
	{"ms", "/majors-suite", "GET", []postData{}, http.StatusOK},
	{"sa", "/search-availability", "GET", []postData{}, http.StatusOK},
	{"contact", "/contact", "GET", []postData{}, http.StatusOK},
	{"reservation-get", "/make-reservation", "GET", []postData{}, http.StatusOK},
	{"post-search-avail", "/search-availability", "POST", []postData{
		{key: "start", value: "2020-01-01"},
		{key: "end", value: "2020-01-02"},
	}, http.StatusOK},
	{"post-search-avail-json", "/search-availability-json", "POST", []postData{
		{key: "start", value: "2020-01-01"},
		{key: "end", value: "2020-01-02"},
	}, http.StatusOK},
	{"make-reservation-start", "/make-reservation", "POST", []postData{
		{key: "first_name", value: "Shamuhammet"},
		{key: "last_name", value: "Ylyasov"},
		{key: "email", value: "shammy@gmail.com"},
		{key: "phone", value: "993622711589"},
	}, http.StatusOK},
}

func TestHandlers(t *testing.T) {
	// getRoutes bize http.Handler return edyar.
	// bu bir yone funksiya. route-leri return edyan
	// setup_test.go-da yazyldy
	routes := getRoutes()

	// httptest.NewTLSServer testowy server start edyar
	// bu dine testler ucin ulanylyar diysek hem bolyar
	// biz her handlerimizi test etmek ucin shona virtualny request ugratmaly bolyarys
	// request ugratmak ucin hem bir server ishlap durmaly.
	// defer ts.Close bolsa hemme requrestler gutaranson start bolan serveri stop edyar
	// httptest.NewTLSServer-in icindaki routes bolsa start edilen serverde gelyan requestleri dinlap durmak ucin
	// onun icinde bizin routes.go-daky routelerimiz bar.
	// routes.go-daky route-leri acylan virtualny serwerde dinlap dur diyen yaly many beryar.
	// request gidende hem sho getRoutes()-den
	ts := httptest.NewTLSServer(routes)
	defer ts.Close()

	for _, e := range theTests {
		if e.method == "GET" {
			// ts.Client bize virtual bir client doredip beryar we Get metody bilen requrest ugradyar
			// bu yerde ts.URL goymagymyzyn sebabi virtualny serverin haysy porta ishlap duranyny bilemzok
			// shonun ucin oz doreden zadyny ozune goyduryarys. shu yerde localhost:8081 yaly bolyar
			// e.url bolsa /, /about bolyar yzygiderlikde.
			// umuman localhost:8081/about (ts.URL/e.url)-a request ugradylyan yaly bolyar
			resp, err := ts.Client().Get(ts.URL + e.url)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		} else {
			//else-in ici POST requestler ucin
			values := url.Values{}
			for _, x := range e.params {
				values.Add(x.key, x.value)
			}

			resp, err := ts.Client().PostForm(ts.URL+e.url, values)
			if err != nil {
				t.Log(err)
				t.Fatal(err)
			}

			if resp.StatusCode != e.expectedStatusCode {
				t.Errorf("for %s expected %d but got %d", e.name, e.expectedStatusCode, resp.StatusCode)
			}
		}
	}
}