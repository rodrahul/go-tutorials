package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

func init() {
	tpl = template.Must(template.ParseGlob("04_templates/*"))
}

func foo(w http.ResponseWriter, req *http.Request) {
	f := req.FormValue("first")
	l := req.FormValue("last")
	s := req.FormValue("subscribe") == "on"

	err := tpl.ExecuteTemplate(w, "index.html",
		person{f, l, s})
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}

func main() {
	http.HandleFunc("/", foo)
	//http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe("localhost:8080", nil)
}
