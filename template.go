package main

import (
	"log"
	"os"
	"text/template"

	"gopkg.in/yaml.v3"
)

// generated from https://zhwt.github.io/yaml-to-go/

type T struct {
	Tactics []struct {
		Description string `yaml:"description"`
		Id          string `yaml:"id"`
		Name        string `yaml:"name"`
		ObjectType  string `yaml:"object-type"`
	} `yaml:"tactics"`
	Techniques []struct {
		AccessRequired      string `yaml:"access-required"`
		ArchitectureSegment string `yaml:"architecture-segment"`
		Bluf                string `yaml:"bluf"`
		Criticalassets      []struct {
			Description string `yaml:"Description"`
			Name        string `yaml:"Name"`
		} `yaml:"criticalassets"`
		Description string `yaml:"description"`
		Detections  []struct {
			Detects string `yaml:"detects"`
			Fgdsid  string `yaml:"fgdsid"`
			Name    string `yaml:"name"`
		} `yaml:"detections"`
		ID          string `yaml:"id"`
		Mitigations []struct {
			Fgmid     string `yaml:"fgmid"`
			Mitigates string `yaml:"mitigates"`
			Name      string `yaml:"name"`
		} `yaml:"mitigations"`
		Name          string `yaml:"name"`
		ObjectType    string `yaml:"object-type"`
		Platforms     string `yaml:"platforms"`
		Preconditions []struct {
			Description string `yaml:"Description"`
			Name        string `yaml:"Name"`
		} `yaml:"preconditions"`
		Procedureexamples []struct {
			Description string `yaml:"Description"`
			Name        string `yaml:"Name"`
		} `yaml:"procedureexamples"`
		References     []string `yaml:"references"`
		Status         string   `yaml:"status"`
		SubtechniqueOf string   `yaml:"subtechnique-of"`
		Tactics        []string `yaml:"tactics"`
		Typecode       string   `yaml:"typecode"`
	} `yaml:"techniques"`
}

type A struct {
	Techniques map[string]string
}

func (t *T) parseFightYaml() *T {
	yamlFile, err := os.ReadFile("fight.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	/*
		for _, technique := range t.Techniques {
			fmt.Printf("%s\n", technique.Name)
		}
	*/

	return t
}

func (a *A) parseAccuknoxYaml() *A {
	yamlFile, err := os.ReadFile("accuknox_support.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, a)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	/*
		for tactic, support := range a.Techniques {
			fmt.Printf("%s: %s\n", tactic, support)
		}
	*/

	return a
}

func (t displayT) generateAllTacticsPage(out *os.File) error {
	tacticList := template.Must(template.New("tacticList").Parse(`
	<h1>Tactics List <h1>
	<table>
	<th style='text-align: left'>
	  {{range .Tactics}}
	  <th><a href="https://5GSEC.github.io/tactic-{{.Id}}.html">{{.Name}}</a></th>
	  {{end}}
	</th>
	</table>
 `))
	if err := tacticList.Execute(out, t); err != nil {
		return err
	}
	return nil
}

func (t Tactic) generateTechniquesPerTacticPage(out *os.File) error {
	techniqueList := template.Must(template.New("techniqueList").Parse(`
	<h1>Tactic:<a href="https://fight.mitre.org/tactics/{{.Id}}">{{.Name}}</a></h1>
	<table>
	<tr style='text-align: left'>
	  <td>Technique Name</td>
	  <td>Accuknox Support</td>
	</th>
	  {{range .Techniques}}
	<tr>
	  <td><a href="https://fight.mitre.org/techniques/{{.Id}}">{{.Name}}</a></td>
	  <td>{{.Support}}</th>
	</tr>
	  {{end}}
	</table>
 `))
	if err := techniqueList.Execute(out, t); err != nil {
		return err
	}
	return nil
}
