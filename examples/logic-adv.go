package main

import (
	_ "embed"
	"os"
	"text/template"
)

// START DATA OMIT
var data = struct {
	Company   string
	Employees []string
}{
	"",
	[]string{"Carson", "Kari", "Tami", "Raul"},
}

// END DATA OMIT

const templateText = `
{{- "" -}}
Company: {{ .Company }}{{ if eq .Company "Weave" }} is great!{{ end }}
{{ if gt (len .Employees) 2 }}Lots of employees!{{ end }}
{{- "" -}}
`

const templateTextView = `
{{- "" -}}
// START TEMPLATE OMIT
Company: {{ .Company }}{{ if eq .Company "Weave" }} is great!{{ end }}
{{ if gt (len .Employees) 2 }}Lots of employees!{{ end }}
// END TEMPLATE OMIT
{{- "" -}}
`

var tmpl = template.Must(template.New("hello").Parse(templateText))

func main() {
	err := tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
		// NOTE: This error is not reachable in this example
	}
}
