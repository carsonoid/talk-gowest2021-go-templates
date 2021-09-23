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
	[]string{"Carson", "Kari", "Tami & Raul"},
}

// END DATA OMIT

const templateText = `
{{- "" -}}
{{ if .Company }}{{ .Company }}{{ else }}No Company{{ end }}
{{ if .Employees }}{{ .Employees }}{{ else }}No Employees{{ end }}
{{- "" -}}
`

const templateTextView = `
{{- "" -}}
// START TEMPLATE OMIT
{{ if .Company }}{{ .Company }}{{ else }}No Company{{ end }}
{{ if .Employees }}{{ .Employees }}{{ else }}No Employees{{ end }}
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
