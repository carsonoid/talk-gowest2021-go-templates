package main

import (
	_ "embed"
	"os"
	"text/template"
)

const templateText = `
{{- "" -}}
Go templates are {{ index .Emojis "fire" }}!
{{ if not (eq .ScaryThing "") }}Just do not get scared by the "{{ .ScaryThing }}" and you will {{ index .Emojis "heart"}} them! {{ end }}
{{- "" -}}
`

const templateTextView = `
{{- "" -}}
// START TEMPLATE OMIT
Go templates are {{ index .Emojis "fire" }}!
{{ if not (eq .ScaryThing "") }}Just do not get scared by the "{{ .ScaryThing }}" and you will {{ index .Emojis "heart"}} them! {{ end }}
// END TEMPLATE OMIT
{{- "" -}}
`

// START CODE OMIT
var data = struct {
	Emojis     map[string]string
	ScaryThing string
}{
	map[string]string{
		"fire":  "🔥",
		"heart": "❤️",
	},
	"syntax",
}

var tmpl = template.Must(template.New("start").Parse(templateText))

// END CODE OMIT

func main() {
	err := tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
		// NOTE: This error is not reachable in this example
	}
}
