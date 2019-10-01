package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Any code you want in this func")
	//fmt.Fprintln(w, req.Body)
	//fmt.Fprintln(w, req.Header)
	fmt.Fprintln(w, req)
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
