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
{{ .Company }}
{{- $msg := "" }}
{{- with len .Employees }}
{{- if eq . 1 }}
{{- $msg = printf "%d is is impossible" . }}
{{- else if eq . 2 }}
{{- $msg = printf "%d is dreary" . }}
{{- else if ge . 3 }}
  {{- $msg = printf "%d is company; safe and cheery" . }}
{{- end -}}
{{- printf " has %03d employees" . }}
{{ printf "%q" $msg }}
{{- end -}}
{{- "" -}}
`

const templateTextView = `
{{- "" -}}
// START TEMPLATE OMIT
{{ .Company }}
{{- $msg := "" }}
{{- with len .Employees }}
{{- if eq . 1 }}
{{- $msg = printf "%d is is impossible" . }}
{{- else if eq . 2 }}
{{- $msg = printf "%d is dreary" . }}
{{- else if ge . 3 }}
  {{- $msg = printf "%d is company; safe and cheery" . }}
{{- end -}}
{{- printf " has %03d employees" . }}
{{ printf "%q" $msg }}
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
