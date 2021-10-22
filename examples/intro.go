package main

import (
	"os"
	"text/template"
)

func main() {
	//START new OMIT
	// Create a new, empty template
	tmpl := template.New("hello")
	//END new OMIT

	//START parse OMIT
	// Parse the template code into the new template
	tmpl, err := tmpl.Parse(`Hello from {{ . }}`)
	// ^ the first returned arg is a pointer to `tmpl` itself
	// so this assignment is essentially pointless here
	if err != nil {
		// NOTE: This error is not reachable in this example
		panic(err)
	}
	//END parse OMIT OMIT

	//START execute OMIT
	// Execute the template with a simple string as data
	err = tmpl.Execute(os.Stdout, "go templates!")
	if err != nil {
		// NOTE: This error is not reachable in this example
		panic(err)
	}
	//END execute OMIT
}
