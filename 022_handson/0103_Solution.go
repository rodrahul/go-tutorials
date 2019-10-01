package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
)

func rootHandleFunc(w http.ResponseWriter, req *http.Request) {
	tpl.ExecuteTemplate(w, "01_03_Solution.html", req.Header)
}

func dogHandleFunc(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Dog says bow bow")
}

func meHandleFunc(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello I'm Rahul Rode")
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("01_03_Solution.html"))
}

func main() {
	http.HandleFunc("/", rootHandleFunc)
	http.HandleFunc("/me/", meHandleFunc)

	http.ListenAndServe(":8080", nil)
}
