package main

import (
	_ "embed"
	"os"
	"text/template"
)

// START DATA OMIT
var data = struct {
	TeamDescriptions map[string]string
	Employees        []string
}{
	map[string]string{"DevX": "Amazing!"},
	[]string{"Carson", "Kari", "Tami", "Raul"},
}

// END DATA OMIT

const templateText = `
{{- "" -}}
Map Range:
{{ range .TeamDescriptions }}{{ println . }}{{ end }}
Slice Range:
{{ range .Employees }}{{ println . }}{{ end }}
{{- "" -}}
`

const templateTextView = `
{{- "" -}}
// START TEMPLATE OMIT
Map Range:
{{ range .TeamDescriptions }}{{ println . }}{{ end }}
Slice Range:
{{ range .Employees }}{{ println . }}{{ end }}
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
