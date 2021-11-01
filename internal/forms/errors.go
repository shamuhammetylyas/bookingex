package forms

type errors map[string][]string

// Add adds an error message for given field
// package level struct. available only inside forms package
// formdan alyan datalarymyzda error bar bolsa sho errory
// shu Add metody bilen errors-a goshuduryarys.
// Son sho errorlar html-daki form inputlaryn ashagynda error hokmunde gorunmeli
func (e errors) Add(field, message string) {
	e[field] = append(e[field], message)
	//e["name"] = ["Name field is required", "Must be at least 3 chars"]
}

// Get, returns first error message
// fieldin errorlaryny almak ucin
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}

	return es[0]
}
