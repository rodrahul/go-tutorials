package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.html"))
}

func rootFunc(w http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(w, "templates/index.html")
	if err != nil {
		log.Fatalln("template didn't execute")
	}

}

func main() {
	http.HandleFunc("/", rootFunc)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("public"))))
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe("localhost:8080", nil)

}
