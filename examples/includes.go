package main

import (
	_ "embed"
	"os"
	"text/template"
)

// START DATA OMIT
type Employee = struct {
	First, Last, Job string
}

var data = struct {
	Company   string
	Employees []Employee
}{
	"Weave",
	[]Employee{
		{"Carson", "Anderson", "Engineer"},
		{"Kari", "Anderson", "Engineer"},
		{"Tami", "Anderson", "Dog"},
	},
}

// END DATA OMIT

const templateText = `
{{- "" -}}
{{- /* sub-template */ -}}
{{- define "employee" }}
{{ .First }} {{ .Last }} - {{ .Job }}
{{- end -}}
Employees:
{{- range .Employees }}
{{- template "employee" . }}
{{- end }}
{{- "" -}}
`

const templateTextView = `
{{- "" -}}
// START TEMPLATE OMIT
{{- /* sub-template */ -}}
{{- define "employee" }}
{{ .First }} {{ .Middle }} {{ .Last }} - {{ .Job }}
{{- end -}}
Employees:
{{- range .Employees }}
{{- template "employee" . }}
{{- end }}
// END TEMPLATE OMIT
{{- "" -}}
`

var tmpl = template.Must(template.New("includes").Parse(templateText))

func main() {
	err := tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
		// NOTE: This error is not reachable in this example
	}
}
