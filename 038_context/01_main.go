package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context()
	//fmt.Println("Context is: ", ctx)
	fmt.Printf("Type of contenxt is %T\n", ctx)
	log.Println(ctx)
	fmt.Fprintln(w, ctx)
}
