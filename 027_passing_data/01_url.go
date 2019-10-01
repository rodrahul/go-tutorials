package main

import (
	"fmt"
	"net/http"
)

func foo(w http.ResponseWriter, req *http.Request) {
	v := req.FormValue("q")
	fmt.Fprintln(w, "Do my search: "+v)
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

//http://localhost:8080/?q=rode
