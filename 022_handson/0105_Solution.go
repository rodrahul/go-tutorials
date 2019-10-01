package main

import (
	"html/template"
	"io"
	"net/http"
)

func rootHandleFunc(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "01_05_Solution.html", req.Header)
}

func dogHandleFunc(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Dog says bow bow")
}

func meHandleFunc(w http.ResponseWriter, req *http.Request) {
	//fmt.Fprintf(w, "Hello I'm Rahul Rode")
	tpl.ExecuteTemplate(w, "01_05_Solution.html", "Rahul Rode")
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("01_05_Solution.html"))
}

func main() {
	http.HandleFunc("/", rootHandleFunc)
	//http.HandleFunc("/me/", meHandleFunc)
	http.Handle("/me/", http.HandlerFunc(meHandleFunc))

	http.ListenAndServe(":8080", nil)
}
