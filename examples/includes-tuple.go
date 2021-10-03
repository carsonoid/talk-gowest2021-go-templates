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
{{- define "employee" }}
{{ .First }} {{ .Middle }} {{ .Last }} - {{ .Job }}
{{- end -}}
Engineers:
{{- range .Employees }}
{{- template "employee" . "Engineers" }}
{{- end }}
Dogs:
{{- range .Employees }}
{{- template "employee" . "Dog" }}
{{- end }}
// END TEMPLATE OMIT
{{- "" -}}
`

const templateTextView = `
{{- "" -}}
// START TEMPLATE OMIT
{{- /* sub-template */ -}}
{{- define "employee" }}
{{ .First }} {{ .Middle }} {{ .Last }} - {{ .Job }}
{{- end -}}

{{- /* reports */ -}}
Engineers: // HL
{{- range .Employees }}
{{- template "employee" . "Engineers" }} // HL
{{- end }}
Dogs: // HL
{{- range .Employees }}
{{- template "employee" . "Dog" }} // HL
{{- end }}
// END TEMPLATE OMIT
{{- "" -}}
`

// START CODE OMIT
// add a `tuple` function for better templates
var tmpl = template.Must(template.New("tuple").
	Funcs(map[string]interface{}{
		"tuple": func(v ...interface{}) []interface{} {
			return v
		},
	}).
	Parse(templateText))

// END CODE OMIT

func main() {
	err := tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
		// NOTE: This error is not reachable in this example
	}
}
