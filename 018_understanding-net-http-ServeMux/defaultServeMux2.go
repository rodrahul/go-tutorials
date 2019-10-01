//ListenAndServe starts an HTTP server with a given address and handler. The handler is usually nil, which means to use DefaultServeMux. Handle and HandleFunc add handlers to DefaultServeMux:
// When using default serve mux you use Hande and HandleFunc

//HandleFunc takes two parameters, first is pattern
//Second is function which has signature (ResponseWriter, *http.Request)
package main

import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {

	//func HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	http.HandleFunc("/dog", d)
	http.HandleFunc("/cat", c)

	http.ListenAndServe(":8080", nil)
}

// Func main can also be written as, if we want to use Handle instead of HandleFunc
/*
func main() {

	// Handle wants route i.e. "/dog" and a handler, (handler is anything which implements ServeHTTP)
	// type HandlerFunc has ServeHTTP method attached to it and type HandlerFunc underlying
	// type is func(ResponseWriter, *Request), which is of type d and c, so now we can use
	// type conversion
	http.Handle("/dog", http.HandlerFunc(d))
	http.Handle("/cat", http.HandlerFunc(c))

	http.ListenAndServe(":8080", nil)

*/
