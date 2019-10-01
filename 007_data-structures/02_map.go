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
	sages := map[string]string{
		"India":    "Gandhi",
		"USA":      "MLK",
		"Medidate": "Buddha",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "02_map.gohtml", sages)
	if err != nil {
		log.Fatal(err)
	}
}
