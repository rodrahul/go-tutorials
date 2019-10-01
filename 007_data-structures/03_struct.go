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
	s := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "03_struct.gohtml", s)
	if err != nil {
		log.Fatal(err)
	}
}
