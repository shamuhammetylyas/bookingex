package forms

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

// Form struct holds form
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

// Post metody gelyan form value-lery alyp bir form doretmek ucin shu New funcstiony yazdyk
// form doretmegimizin peydasy ashakdaky receiver metodlary ulanyp bilyas son.
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
func (f *Form) Has(field string) bool {
	// x := req.Form.Get(field)
	// f.Get => shu yerdaki Get funcksiya form pointerin icindaki url.Values-in get metodydyr. Embedded struct ulanylany ucin sheyle yazylyar
	x := f.Get(field)
	return x != ""
	// if x == "" {
	// 	return false
	// }
	// return true
}

func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		// f.Get => shu yerdaki Get funcksiya form pointerin icindaki url.Values-in get metodydyr. Embedded struct ulanylany ucin sheyle yazylyar
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

func (f *Form) MinLength(field string, length int8) bool {
	// f.Get => shu yerdaki Get funcksiya form pointerin icindaki url.Values-in get metodydyr. Embedded struct ulanylany ucin sheyle yazylyar
	x := f.Get(field)
	if len(x) < int(length) {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long", length))
		return false
	}

	return true
}

func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(field) {
		f.Errors.Add(field, "Invalid Email address")
	}
}
