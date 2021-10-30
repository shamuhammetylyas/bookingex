package forms

import (
	"net/http"
	"net/url"
)

//Form struct holds form
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

// Has checks if form field is in POST and not empty
func (f *Form) Has(field string, req *http.Request) bool {
	x := req.Form.Get(field)
	return x != ""
	// if x == "" {
	// 	return false
	// }

	// return true
}
