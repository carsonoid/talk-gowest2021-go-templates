package main

import (
	"os"
	"text/template"
)

const (
	hello   = `Hello from {{ . }}`
	primary = `My sub-template says: {{ template "hello" . }}`
)

func main() {
	tmpl, err := template.New("hello").Parse(hello)
	if err != nil {
		panic(err)
	}

	tmpl, err = tmpl.New("primary").Parse(primary)
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, "go templates!")
	if err != nil {
		panic(err)
	}
}
