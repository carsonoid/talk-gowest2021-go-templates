package main

import (
	"os"
	"text/template"
)

func main() {
	tmpl, _ := template.New("hello").Parse(`Hello from {{ . }}`)
	// ^ this is why .Parse returns a Template. This more common and
	//   compact usage requires that parse return it so we can use it

	_ = tmpl.Execute(os.Stdout, "go templates!")
}
