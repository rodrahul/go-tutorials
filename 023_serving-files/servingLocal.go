package main

import (
	"io"
	"net/http"
	"os"
)

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
		<img src="/_DSC0132-Edit-copy.jpg">
		`)
}

func dscPic(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("_DSC0132-Edit-copy.jpg")
	if err != nil {
		http.Error(w, "Picture not found", 404)
		return
	}
	defer f.Close()

	io.Copy(w, f)
}

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/_DSC0132-Edit-copy.jpg", dscPic)

	http.ListenAndServe(":8080", nil)
}
