package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("01_data.gohtml"))
}
func main(){
	err := tpl.ExecuteTemplate(os.Stdout, "01_data.gohtml", 42)
	if err != nil {
		log.Fatal(err)
	}
}
