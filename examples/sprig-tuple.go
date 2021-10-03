// START CODE OMIT
// use the full sprig funcmap
var tmpl = template.Must(template.New("tuple").
	Funcs(sprig.FuncMap()).
	Parse(templateText))

// END CODE OMIT
