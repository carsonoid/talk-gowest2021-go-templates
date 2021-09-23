package main

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"text/template"
)

//go:embed templates
var templateFS embed.FS
var tmpl = template.Must(template.New("t").ParseFS(templateFS, "templates/*"))

func main() {
	files, _ := fs.ReadDir(templateFS, "templates")
	fmt.Println("Templates:")
	for _, file := range files {
		fmt.Println(file.Name())
	}
	fmt.Println("\nExecute hello.tmpl:")
	_ = tmpl.ExecuteTemplate(os.Stdout, "hello.tmpl", "go templates!")
}
