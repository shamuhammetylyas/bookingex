package forms

import (
	"net/http"
	"net/url"
)

//Form struct holds form
// form structymyz bosh bir html acylanda-da gerek bolyar
// form submit edilenson errorlary html formda gorkezmek ucin hem gerek bolyar
// bu form hem formdan gelyan datalary saklap bilyar, hem formyn errolaryny saklayar
// form sahypasy birinji gezek render edilende formda hic hili error bolanok
// shon ucin New form doredilende data hokmunde nil ugradylyar.
// eger form submit edilenson onda hem error bar bolsa onda data hokmunde onki yazylan
// form valuelar we doredilen errorlar form bilen ugradylyar
type Form struct {
	url.Values
	Errors errors
}

func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// valid return true if there are no errors, otherwise false
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// Has checks if form field is in POST and not empty
func (f *Form) Has(field string, req *http.Request) bool {
	x := req.Form.Get(field)
	// return x != ""
	if x == "" {
		f.Errors.Add(field, "This field cannot be blank")
		return false
	}

	return true
}
