package main

import (
	_ "embed"
	"fmt"
	"os"
	"text/template"
)

const templateText = `
{{- "" -}}
Go templates are {{ emoji "fire" }}!
{{- with .ScaryThing }}
Just do not get scared by the {{ . | quote }} and you will {{ emoji "heart" }} them!
{{- end }}
{{- "" -}}
`

const templateTextView = `
{{- "" -}}
// START TEMPLATE OMIT
Go templates are {{ emoji "fire" }}!
{{- with .ScaryThing }}
Just do not get scared by the {{ . | quote }} and you will {{ emoji "heart" }} them!
{{- end }}
// END TEMPLATE OMIT
{{- "" -}}
`

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

// START CODE OMIT
// create a template with some helper funcs
var tmpl = template.Must(template.New("start").Funcs(
	map[string]interface{}{
		// closure on data.emojis
		"emoji": func(name string) string {
			if emoji, ok := data.Emojis[name]; ok {
				return emoji
			}
			return "❓"
		},
		// easy quoter func
		"quote": func(s string) string {
			return fmt.Sprintf("%q", s)
		},
	},
).Parse(templateText))

// END CODE OMIT

func main() {
	err := tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
		// NOTE: This error is not reachable in this example
	}
}
