package main

import (
	"fmt"
	"io"
	"net/http"
)

func rootHandleFunc(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "You are at the root of the server")
}

func dogHandleFunc(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Dog says bow bow")
}

func meHandleFunc(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello I'm Rahul Rode")
}

func main() {
	http.HandleFunc("/", rootHandleFunc)
	http.HandleFunc("/dog/", dogHandleFunc)
	http.HandleFunc("/me/", meHandleFunc)

	http.ListenAndServe(":8080", nil)
}
