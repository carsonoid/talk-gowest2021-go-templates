package main

import (
	"os"
	"text/template"
)

// go:embed templates/hello.tmpl
var helloText string

// go:embed templates/sub.tmpl
var subTemplateText string

func main() {
	tmpl, err := template.New("hello").Parse(helloText)
	if err != nil {
		panic(err)
	}

	tmpl, err = tmpl.New("adv").
		Parse(subTemplateText)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, "go templates!")
	if err != nil {
		panic(err)
	}
}
