package main

import (
	_ "embed"
	"os"
	"text/template"
)

//go:embed templates/uglier-fixed.go.tmpl
var tmplText string

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
	helpers := map[string]interface{}{
		"employeeTeams": employeeTeams,
	}

	template.Must(template.New("").
		Funcs(helpers).
		Parse(tmplText)).
		Execute(os.Stdout, data)
	// END CODE OMIT
}
