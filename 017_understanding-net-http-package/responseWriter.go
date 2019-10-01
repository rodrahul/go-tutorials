package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Mcleod-Key", "this is from mcleod")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>Any code you want in this func</h1>\n", m)
	//w.Header().Write(w)
	//fmt.Fprintln(w, w.Header())

}

func main() {
	//var d hotdog = 10
	d := hotdog(10)
	var i int
	i = int(d)
	fmt.Printf("Type of d is %t and i is %t\n", d, i)
	http.ListenAndServe(":8080", d)
}
