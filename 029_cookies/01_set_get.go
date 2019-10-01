package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func set(w http.ResponseWriter, req *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "rahul-cookie",
		Value: "some-value",
		Path:  "/",
	})

	fmt.Fprintf(w, "COOKIE WRITTEN - CHECK YOUR BROWSER")
	fmt.Fprintln(w, "in chrome go to: dev tools / application / cookies")
}

func read(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("rahul-cookie")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w, "Cookie Read is ", cookie)
}
