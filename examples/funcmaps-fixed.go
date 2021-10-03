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
{{- $msg := getMessage $num }} // HL
{{ printf "%s has %d employees" .Company $num  }}
{{ printf "%d %s" $num $msg | printf "> %q" }}
// END TEMPLATE OMIT
{{- "" -}}
`

// START CODE OMIT
func getMessage(num int) string {
	switch num {
	case 1:
		return "is impossbile"
	case 2:
		return "is dreary"
	case 3:
		return "is company; safe and cheery"
	}
	return ""
}

func main() {
	funcMap := map[string]interface{}{
		"getMessage": getMessage,
	}

	tmpl := template.Must(
		template.New("variables").
			Funcs(funcMap).
			Parse(templateText),
	)
	_ = tmpl.Execute(os.Stdout, data)
}

// END CODE OMIT
