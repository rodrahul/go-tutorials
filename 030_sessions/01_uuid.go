package main

import (
	"github.com/satori/go.uuid"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/read", read)

	http.ListenAndServe("localhost:8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	uuid, _ := uuid.NewV4()
	cookie := &http.Cookie{
		Name:     "session",
		Value:    uuid.String(),
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
	<p>UUID was written to the cookie check your browser</p>
	<br>
	<a href=/read >read
	`)
}

func read(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("session")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	io.WriteString(w, `
	<p>Reading cookie from browser storage</p>`+cookie.Value)
	io.WriteString(w, `<br>
	<a href=/foo >Set New Cookie
	`)
}
