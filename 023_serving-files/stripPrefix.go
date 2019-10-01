package main

import (
	"io"
	"net/http"
)

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/resources/cathedral.jpg">`)
}

func main() {
	http.HandleFunc("/", dog)
	http.Handle("/resources/", http.StripPrefix("/resources", http.FileServer(http.Dir("./assets"))))
	http.ListenAndServe(":8080", nil)
}

// To serve a directory on disk (/tmp) under an alternate URL
// path (/tmpfiles/), use StripPrefix to modify the request
// URL's path before the FileServer sees it:
// http.Handle("/tmpfiles/", http.StripPrefix("/tmpfiles/", http.FileServer(http.Dir("/tmp"))))
