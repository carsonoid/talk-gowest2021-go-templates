package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)

var tmpl = template.Must(template.New("t").ParseGlob("templates/*"))

func main() {
	files, _ := ioutil.ReadDir("templates")
	fmt.Println("Templates:")
	for _, file := range files {
		fmt.Println(file.Name())
	}

	fmt.Println("\nExecute hello.tmpl:")
	_ = tmpl.ExecuteTemplate(os.Stdout, "hello.tmpl", "go templates!")
}
