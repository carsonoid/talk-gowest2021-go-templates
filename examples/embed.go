package main

import (
	_ "embed"
	"os"
	"text/template"
)

//go:embed templates/hello.tmpl
var templateText string

var tmpl = template.Must(template.New("hello").Parse(templateText))

func main() {
	err := tmpl.Execute(os.Stdout, "go templates!")
	if err != nil {
		panic(err)
		// NOTE: This error is not reachable in this example
	}
}
