package main

import (
	html "html/template"
	"os"
	text "text/template"
)

var htmlTmpl = html.Must(html.New("hello").Parse("Hello from {{ . }}\n"))
var textTmpl = text.Must(text.New("hello").Parse("Hello from {{ . }}\n"))

func main() {
	htmlTmpl.Execute(os.Stdout, "--> html & the template package <---")
	textTmpl.Execute(os.Stdout, "--> text & the template package <---")
}
