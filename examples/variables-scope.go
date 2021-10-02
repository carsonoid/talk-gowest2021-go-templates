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
{{- if eq $num 1 }}
  {{- $msg := "is imposible" }}
{{- else if eq $num 2 }}
  {{- $msg := "is dreary" }}
{{- else if ge $num 3 }}
  {{- $msg := "is company; safe and cheery" }}
{{- else }}
  {{- $msg := "I have no comment" }}
{{- end -}}

{{ printf " has employees: %03d %s" $num $msg }}
{{- "" -}}
`

const templateTextView = `
{{- "" -}}
// START TEMPLATE OMIT
{{ .Company }}:

{{- $num := len .Employees }}
{{- if eq $num 1 }}
  {{- $msg := "is imposible" }}
{{- else if eq $num 2 }}
  {{- $msg := "is dreary" }}
{{- else if ge $num 3 }}
  {{- $msg := "is company; safe and cheery" }}
{{- else }}
  {{- $msg := "I have no comment" }}
{{- end -}}

{{ printf " has employees: %03d %s" $num $msg }}
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
