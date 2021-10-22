package main

import (
	"os"
	"text/template"
)

const templateText = `Hello from {{ . }}
Here is another line.
Hello again from {{ . }}
`

var tmpl = template.Must(template.New("hello").Parse(templateText))

func main() {
	err := tmpl.Execute(os.Stdout, "go templates!")
	if err != nil {
		panic(err)
		// NOTE: This error is not reachable in this example
	}
}
