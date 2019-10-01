package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseGlob("./template/*"))
}

func main() {
	b := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}
	g := sage{
		Name:  "Gandhi",
		Motto: "Be the change",
	}
	mlk := sage{
		Name:  "MLK",
		Motto: "Hatred never ceases with hatred but with love alone is healed",
	}
	sageSlice := []sage{
		b,g,mlk,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "04_slice-struct.gohtml", sageSlice)
	if err != nil {
		log.Fatal(err)
	}
}
