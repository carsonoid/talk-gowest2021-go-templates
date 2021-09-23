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
	[]string{"Carson", "Kari", "Tami", "Raul"},
}

// END DATA OMIT

const templateText = `
{{- "" -}}
Company: {{ .Company }}
{{- /* A special note for Weave :) */}}
{{- if eq .Company "Weave" }} is great!{{ end }}
{{/* comments can also not chomp */}}
{{- if gt (len .Employees) 2 }}
Lots of employees!
{{- end }}
{{- "" -}}
`

const templateTextView = `
{{- "" -}}
// START TEMPLATE OMIT
Company: {{ .Company }}
{{- /* A special note for Weave :) */}}
{{- if eq .Company "Weave" }} is great!{{ end }}
{{/* comments can also not chomp */}}
{{- if gt (len .Employees) 2 }}
Lots of employees!
{{- end }}
// END TEMPLATE OMIT
{{- "" -}}
`

var tmpl = template.Must(template.New("comments").Parse(templateText))

func main() {
	err := tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
		// NOTE: This error is not reachable in this example
	}
}
