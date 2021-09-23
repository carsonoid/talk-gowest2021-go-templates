package main

import "fmt"

func main() {

	// START OMIT
	var data = struct {
		Company   string
		Employees []string
	}{
		"Weave",
		[]string{"Carson", "Kari", "Tami", "Raul"},
	}

	fmt.Println("Company:", data.Company)
	if len(data.Employees) > 0 {
		fmt.Println("Employees")
		for _, employee := range data.Employees {
			fmt.Println("-", employee)
		}
	}
	// END OMIT
}
