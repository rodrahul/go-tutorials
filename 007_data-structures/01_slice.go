package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("./template/*"))
}

func main() {
	sages := []string{
		"Gandhi",
		"MKL",
		"Buddha",
		"Jesus",
		"Muhammad",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "01_slice.gohtml", sages)
	if err != nil {
		log.Fatal(err)
	}
}
