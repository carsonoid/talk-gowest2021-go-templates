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
{{- $num := len .Employees }}
{{ .Company }}{{ printf " has %d employees" $num }}
{{- if eq $num 1 }}
{{ printf "%d is impossible" $num | printf "> %q" }}
{{- else if eq $num 2 }}
{{ printf "%d is dreary" $num | printf "> %q" }}
{{- else if ge $num 3 }}
{{ printf "%d is company; safe and cheery" $num | printf "> %q" }}
{{- end -}}
{{- "" -}}
`

const templateTextView = `
{{- "" -}}
// START TEMPLATE OMIT
-Company Report-
{{- $num := len .Employees }}
{{ .Company }}{{ printf " has %d employees" $num }}
{{- if eq $num 1 }}
{{ printf "%d is impossible" $num | printf "> %q" }}
{{- else if eq $num 2 }}
{{ printf "%d is dreary" $num | printf "> %q" }}
{{- else if ge $num 3 }}
{{ printf "%d is company; safe and cheery" $num | printf "> %q" }}
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
