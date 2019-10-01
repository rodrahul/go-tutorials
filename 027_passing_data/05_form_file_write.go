package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("04_templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	var s string
	if req.Method == http.MethodPost {

		f, fh, err := req.FormFile("q")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		fmt.Println("Filename:", fh.Filename,
			"File header:", fh.Header,
			"File Size:", fh.Size)
		bs, _ := ioutil.ReadAll(f)
		s = string(bs)
		//File read finished
		// now write the file
		// Create the file with the same name as the uploaded file, and store it in user dir
		dest, _ := os.Create(filepath.Join("./user/", fh.Filename))
		fmt.Printf("\nType of f is %T and of dest is %T\n", f, dest)
		dest.Write(bs)

	}

	w.Header().Set("Content-Type", "text/html charset=utf-8")
	tpl.ExecuteTemplate(w, "05_form_file_write_index.html", s)
}
