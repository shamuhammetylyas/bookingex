package models

import "github.com/ShamuhammetYlyas/bookings/internal/forms"

//TemplateData holds data sent from handlers to templates
// shu template data her gezek sahypa render edilende shu datalar hem gidyar render edilen sahypa
// bosh gidyan bolmagam mumkin value-ly gidyan bolmagam mumkin
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
	Form      *forms.Form
}
