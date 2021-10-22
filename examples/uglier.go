package main

import "fmt"

func main() {
	// START DATA OMIT
	var data = struct {
		Company   string
		Employees []string
		Teams     map[string][]string
	}{
		"Weave",
		[]string{"Carson", "Kari", "Tami", "Raul"},
		map[string][]string{
			"Leads":   {"Carson", "Kari"},
			"Interns": {"Tami", "Raul"},
		},
	}
	// END DATA OMIT

	employeeTeams := func(employee string) []string {
		var teams []string
		for team, members := range data.Teams {
			for _, member := range members {
				if employee == member {
					teams = append(teams, team)
				}
			}
		}
		return teams
	}

	// START CODE OMIT
	fmt.Println("Company:", data.Company)
	if len(data.Employees) > 0 {
		fmt.Println("Employees")
		for _, employee := range data.Employees {
			fmt.Println("-", employee)
			teams := employeeTeams(employee)
			if len(teams) > 0 {
				fmt.Println("  Teams:")
				for _, team := range teams {
					fmt.Println("    -", team)
				}
			}
		}
	}
	// END CODE OMIT
}
