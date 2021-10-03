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
	Features  map[string]bool
}{
	"Weave",
	[]string{"Carson", "Kari", "Tami & Raul"},
	map[string]bool{
		"beta-db": true,
		"new-ui":  false,
	},
}

// END DATA OMIT

const templateText = `
{{- "" -}}
{{ .Company }} Status:
New UI:      {{ feature "new-ui" }}
Beta DB:     {{ feature "beta-db" }}
Auto Update: {{ feature "auto-update" }}
{{- "" -}}
`

const templateTextView = `
{{- "" -}}
// START TEMPLATE OMIT
{{ .Company }} Status:
New UI:      {{ feature "new-ui" }}
Beta DB:     {{ feature "beta-db" }}
Auto Update: {{ feature "auto-update" }}
// END TEMPLATE OMIT
{{- "" -}}
`

func main() {
	// START CODE OMIT
	// var data = ...

	feature := func(name string) string {
		if val, ok := data.Features[name]; ok {
			if val {
				return "enabled"
			}
			return "disabled"
		}
		return "unset"
	}

	funcMap := map[string]interface{}{
		"feature":    feature,
		"getMessage": getMessage,
	}
	// END CODE OMIT

	tmpl := template.Must(
		template.New("variables").Funcs(funcMap).Parse(templateText),
	)
	_ = tmpl.Execute(os.Stdout, data)
}

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
