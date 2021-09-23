package main

import (
	_ "embed"
	"os"
	"text/template"
	"time"
)

var tmpl = template.Must(template.New("hello").Parse(templateText))

// START DATA OMIT
var data = struct {
	EventChan chan string
}{
	make(chan string),
}

func main() {
	go func() {
		for range time.Tick(time.Second) {
			data.EventChan <- "hello!"
		}
	}()
	go func() { time.Sleep(time.Second * 4); close(data.EventChan) }()
	_ = tmpl.Execute(os.Stdout, data)
}

// END DATA OMIT

const templateText = `
{{- "" -}}
// START TEMPLATE OMIT
Events:
{{ range .EventChan }}{{ println .  }}{{ end }}
// END TEMPLATE OMIT
{{- "" -}}
`
