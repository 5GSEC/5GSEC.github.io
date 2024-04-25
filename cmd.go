package main

import (
	"fmt"
	"os"
)

type Tactic struct {
	Id         string
	Name       string
	Techniques []Technique
}

type Technique struct {
	Name    string
	Id      string
	Support string
}

type displayT struct {
	Tactics []Tactic
}

func main() {

	// parse the yaml
	var t T
	t = *t.parseFightYaml()
	var a A
	a = *a.parseAccuknoxYaml()

	// copy the support from the accuknox support yaml

	// populate the display var which is input to template renderer
	var d displayT
	for _, tactic := range t.Tactics {
		var t_d Tactic
		t_d.Id = tactic.Id
		t_d.Name = tactic.Name
		for _, technique := range t.Techniques {
			for _, tqt := range technique.Tactics {
				if t_d.Id == tqt {
					t_d.Techniques = append(t_d.Techniques, Technique{technique.Name, technique.ID, a.Techniques[technique.Name]})
				}
			}
		}
		d.Tactics = append(d.Tactics, t_d)
	}

	// print the struct
	for _, tactic := range d.Tactics {
		for _, technique := range tactic.Techniques {
			fmt.Printf("%s: %s: %s\n", tactic.Name, technique.Name, technique.Support)
		}
	}

	// open output file
	fo, err := os.Create("index.html")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := fo.Close(); err != nil {
			panic(err)
		}
	}()

	// render the template
	d.generateAllTacticsPage(fo)

	for _, tactic := range d.Tactics {

		// open output file
		fname := fmt.Sprintf("tactic-%s.html", tactic.Id)
		fo, err := os.Create(fname)
		if err != nil {
			panic(err)
		}

		defer func() {
			if err := fo.Close(); err != nil {
				panic(err)
			}
		}()

		// render the template
		tactic.generateTechniquesPerTacticPage(fo)
	}
}
