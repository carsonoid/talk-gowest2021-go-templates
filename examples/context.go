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
	"Weave",
	[]string{"Carson", "Kari", "Tami"},
}

// END DATA OMIT

const templateText = `
{{- "" -}}
{{- .Company }}:
{{- if .Employees }}
Employees: {{ .Employees }}
{{- end }}
{{- "" -}}
`

const templateTextView = `
// START TEMPLATE OMIT
{{- .Company }}:
{{- if .Employees }}
Employees: {{ .Employees }}
{{- end }}
// END TEMPLATE OMIT
`

var tmpl = template.Must(template.New("hello").Parse(templateText))

func main() {
	err := tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
		// NOTE: This error is not reachable in this example
	}
}
