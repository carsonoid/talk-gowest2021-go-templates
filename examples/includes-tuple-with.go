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
{{- $job := index . 1 }}
{{- with index . 0}}
{{- if eq $job .Job }}
{{ .First }} {{ .Last }} - {{ .Job }}
{{- end -}}
{{- end -}}
{{- end -}}

{{- /* reports */ -}}
Engineers:
{{- range .Employees }}
{{- template "employee" (tuple . "Engineer") }}
{{- end }}
Dogs:
{{- range .Employees }}
{{- template "employee" (tuple . "Dog") }}
{{- end }}
{{- "" -}}
`

const templateTextView = `
{{- "" -}}
// START TEMPLATE OMIT
{{- /* sub-template */ -}}
{{- define "employee" }}
{{- $job := index . 1 }}
{{- with index . 0}}
{{- if eq $job .Job }}
{{ .First }} {{ .Last }} - {{ .Job }}
{{- end -}}
{{- end -}}

{{- /* reports */ -}}
Engineers:
{{- range .Employees }}
{{- template "employee" (tuple . "Engineer") }}
{{- end }}
Dogs:
{{- range .Employees }}
{{- template "employee" (tuple . "Dog") }}
{{- end }}
// END TEMPLATE OMIT
{{- "" -}}
`

var tmpl = template.Must(template.New("includes").
	Funcs(map[string]interface{}{
		"tuple": func(v ...interface{}) []interface{} {
			return v
		},
	}).
	Parse(templateText))

func main() {
	err := tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
		// NOTE: This error is not reachable in this example
	}
}
