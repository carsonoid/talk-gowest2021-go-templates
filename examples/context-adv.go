package main

import (
	_ "embed"
	"os"
	"text/template"
)

const templateText = `
{{- "" -}}
My context is:                  {{ . }}

{{- /* You can use variables to store data before a context change */}}
{{- with .Company }}
"with .Company" the context is: {{ . }}
{{- end }}
{{- "" -}}
`

const templateTextView = `
// START TEMPLATE OMIT
My context is:                  {{ . }}

{{- /* You can use variables to store data before a context change */}}
{{- with .Company }}
"with .Company" the context is: {{ . }}
{{- end }}
// END TEMPLATE OMIT
`

func main() {
	// START CODE OMIT
	var data = struct {
		Company   string
		Employees []string
	}{
		"Weave",
		[]string{"Carson", "Kari", "Tami"},
	}

	tmpl := template.Must(template.New("context").Parse(templateText))
	_ = tmpl.Execute(os.Stdout, data)
	// END CODE OMIT
}
