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
{{ .Company }}:
{{- $num := len .Employees }}
{{ printf "# employees: %03d" $num }}
{{- if eq $num 1 }} is imposible
{{- else if eq $num 2 }} is dreary
{{- else if ge $num 3 }} is company; safe and cheery
{{- else }}- I have no comment{{- end }}
{{- "" -}}
`

const templateTextView = `
{{- "" -}}
// START TEMPLATE OMIT
{{ .Company }}:
{{ $num := len .Employees }}
{{ printf "# employees: %03d" $num }}
{{- if eq $num 1 }} is imposible
{{- else if eq $num 2 }} is dreary
{{- else if ge $num 3 }} is company; safe and cheery
{{- else }}- I have no comment{{- end }}
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
