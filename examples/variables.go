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
	[]string{"Carson", "Kari", "Tami & Raul"},
}

// END DATA OMIT

const templateText = `
{{- "" -}}
-Company Report-
{{ .Company }}{{ printf " has %d employees" (len .Employees) }}
{{- if eq (len .Employees) 1 }}
{{ printf "%d is impossible" (len .Employees) | printf "> %q" }}
{{- else if eq (len .Employees) 2 }}
{{ printf "%d is dreary" (len .Employees) | printf "> %q" }}
{{- else if ge (len .Employees) 3 }}
{{ printf "%d is company; safe and cheery" (len .Employees) | printf "> %q" }}
{{- end -}}
{{- "" -}}
`

const templateTextView = `
{{- "" -}}
// START TEMPLATE OMIT
-Company Report-
{{ .Company }}{{ printf " has %d employees" (len .Employees) }}
{{- if eq (len .Employees) 1 }}
{{ printf "%d is impossible" (len .Employees) | printf "> %q" }}
{{- else if eq (len .Employees) 2 }}
{{ printf "%d is dreary" (len .Employees) | printf "> %q" }}
{{- else if ge (len .Employees) 3 }}
{{ printf "%d is company; safe and cheery" (len .Employees) | printf "> %q" }}
{{- end -}}
// END TEMPLATE OMIT
{{- "" -}}
`

var tmpl = template.Must(template.New("variables").Parse(templateText))

func main() {
	err := tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
		// NOTE: This error is not reachable in this example
	}
}
