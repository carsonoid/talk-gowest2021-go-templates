package main

import (
	_ "embed"
	"fmt"
	"text/template"
)

// START DATA OMIT
var data = struct {
	TeamDescriptions map[string]string
	Employees        []string
}{
	map[string]string{"DevX": "Amazing!"},
	[]string{"Carson", "Kari", "Tami & Raul"},
}

// END DATA OMIT

const templateText = `
{{- "" -}}
Team Descriptions:
{{ range .TeamDescriptions }}{{ println . }}{{ end }}
Employees:
{{ range .Employees }}{{ println . }}{{ end }}
{{- "" -}}
`

const templateTextView = `
{{- "" -}}

{{- "" -}}
`

var tmpl = template.Must(template.New("hello").Parse(templateText))

func main() {
	// START CODE OMIT
	fmt.Println("Map Range:")
	for x := range data.TeamDescriptions {
		fmt.Println(x)
	}
	fmt.Println("Slice Range:")
	for x := range data.Employees {
		fmt.Println(x)
	}
	// END CODE OMIT
}
