package main

import (
	"io/ioutil"
	"os"
	"text/template"
)

func main() {
	// read file
	templateText, _ := ioutil.ReadFile("templates/hello.tmpl")

	// create template
	tmpl := template.Must(template.New("hello").Parse(string(templateText)))

	// execute
	_ = tmpl.Execute(os.Stdout, "go templates!")
}
