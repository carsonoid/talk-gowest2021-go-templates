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
	[]string{"Carson", "Kari", "Tami"},
}

// END DATA OMIT

const templateText = `
{{- "" -}}
-Company Report-
{{- $num := len .Employees }}
{{- $msg := "" }}
{{- if eq $num 1 }}
{{- $msg = "is impossible" }}
{{- else if eq $num 2 }}
{{- $msg = "is dreary" }}
{{- else if ge $num 3 }}
{{- $msg = "is company; safe and cheery" }}
{{- end }}
{{ printf "%s has %d employees" .Company $num  }}
{{ printf "%d %s" $num $msg | printf "> %q" }}
{{- "" -}}
`

const templateTextView = `
{{- "" -}}
// START TEMPLATE OMIT
-Company Report-
{{- $num := len .Employees }}
{{- $msg := "" }} // HL
{{- if eq $num 1 }}
{{- $msg = "is impossible" }}
{{- else if eq $num 2 }}
{{- $msg = "is dreary" }}
{{- else if ge $num 3 }}
{{- $msg = "is company; safe and cheery" }}
{{- end }}
{{ printf "%s has %d employees" .Company $num  }}
{{ printf "%d %s" $num $msg | printf "> %q" }}
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
