package main

import (
	"os"
	"text/template"
)

func main() {
	tmpl, err := template.New("hello").Parse(`Hello from {{ . }}`)
	if err != nil {
		panic(err)
	}

	tmpl, err = tmpl.New("adv").
		Parse(`My sub-template says: {{ template "hello" . }}`)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, "go templates!")
	if err != nil {
		panic(err)
	}
}
