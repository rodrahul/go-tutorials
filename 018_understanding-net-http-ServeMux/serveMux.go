package main

import (
	"io"
	"net/http"
)

type hotdog int

func (d hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Dog Dog Dog")
}

type hotcat int

func (c hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "Cat Cat Cat")
}

func main() {
	var d hotdog
	var c hotcat
	mux := http.NewServeMux()
	mux.Handle("/dog/", d)
	mux.Handle("/cat", c)

	http.ListenAndServe(":8080", mux)
}
