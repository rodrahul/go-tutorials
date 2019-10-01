package main

import (
	"fmt"
	"io"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	_, _ = fmt.Fprintf(res, "url path is %s\n\n", req.URL.Path)

	switch req.URL.Path {

	case "/dog":
		_, _ = io.WriteString(res, "doggy doggy")
	case "/cat":
		_, _ = io.WriteString(res, "meow meow")
	default:
		_, _ = io.WriteString(res, "default")
	}
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
