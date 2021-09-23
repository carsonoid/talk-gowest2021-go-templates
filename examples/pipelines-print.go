package main

import (
	_ "embed"
	"fmt"
)

// START DATA OMIT
var data = struct {
	Company   string
	Employees []string
}{
	"Weave",
	[]string{"Carson", "Kari", "Jake"},
}

// END DATA OMIT

// Data: {{ . }}
// Company: {{ .Company }}
// Employees: {{ .Employees }}

// START TEMPLATE OMIT
func main() {
	fmt.Print("Data:      ", data, "\n")
	fmt.Print("Company:   ", data.Company, "\n")
	fmt.Print("Employees: ", data.Employees)
}

// END TEMPLATE OMIT
