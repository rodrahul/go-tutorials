package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

var count int

func main() {
	http.HandleFunc("/", setGet)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe(":8080", nil)
}

func setGet(w http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("visit-counter")
	if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "visit-counter",
			Value: "0",
			Path:  "/",
		}
		fmt.Println("Setting cookie")
		fmt.Println("Error is ", err.Error())
	}
	count, _ = strconv.Atoi(cookie.Value)
	count++
	cookie.Value = strconv.Itoa(count)
	http.SetCookie(w, cookie)

	io.WriteString(w, cookie.Value)

}
