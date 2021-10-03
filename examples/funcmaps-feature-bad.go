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
	// START FEATS OMIT
	map[string]bool{
		"beta-db": true,
		"new-ui":  false,
	},
	// END FEATS OMIT
}

// END DATA OMIT

const templateText = `
{{- "" -}}
{{ .Company }} Status:
New UI:      {{ if index .Features "new-ui" }}enabled{{ else }}disabled{{ end }}
Beta DB:     {{ if index .Features "beta-db" }}enabled{{ else }}disabled{{ end }}
Auto Update: {{ if index .Features "auto-update" }}enabled{{ else }}disabled{{ end }}
{{- "" -}}
`

const templateTextView = `
{{- "" -}}
// START TEMPLATE OMIT
{{ .Company }} Status:
New UI:      {{ if index .Features "new-ui" }}enabled{{ else }}disabled{{ end }}
Beta DB:     {{ if index .Features "beta-db" }}enabled{{ else }}disabled{{ end }}
Auto Update: {{ if index .Features "auto-update" }}enabled{{ else }}disabled{{ end }}
// END TEMPLATE OMIT
{{- "" -}}
`

func main() {
	// START CODE OMIT
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
