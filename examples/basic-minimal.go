package main

import (
	"os"
	"text/template"
)

var tmpl = template.Must(template.New("hello").Parse(`Hello from {{ . }}`))

func main() {
	err := tmpl.Execute(os.Stdout, "go templates!")
	if err != nil {
		panic(err)
		// NOTE: This error is not reachable in this example
	}
}
