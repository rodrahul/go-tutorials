package main

import (
	"crypto/sha1"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.Handle("/public/", http.StripPrefix("/public", http.FileServer(http.Dir("./public"))))
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	c := getCookie(w, req)

	if req.Method == http.MethodPost {
		f, fh, err := req.FormFile("nf")
		if err != nil {
			http.Error(w, "failed to read file", http.StatusInternalServerError)
			return
		}
		defer f.Close()

		// create sha for file name
		ext := strings.Split(fh.Filename, ".")[1]
		h := sha1.New()
		io.Copy(h, f)
		fname := fmt.Sprintf("%x", h.Sum(nil)) + "." + ext

		//Create new file
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}
		path := filepath.Join(wd, "public", "pics", fname)
		nf, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
		}
		defer nf.Close()

		f.Seek(0, 0)
		io.Copy(nf, f)
		fmt.Println("file name is", fname)

		// add file name to the users cookie
		c = appendCookie(w, c, fname)
	}
	xs := strings.Split(c.Value, "|")
	// sliced cookie values to only send over images
	tpl.ExecuteTemplate(w, "index.html", xs[1:])

}

func appendCookie(w http.ResponseWriter, c *http.Cookie, fname string) *http.Cookie {
	s := c.Value
	if !strings.Contains(s, fname) {
		s += "|" + fname
	}
	c.Value = s
	http.SetCookie(w, c)
	return c
}

func getCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {
	c, err := req.Cookie("session")
	if err != nil {
		fmt.Println("Cookie not found, setting new cookie, err is ", err)
		sId, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sId.String(),
		}
		http.SetCookie(w, c)
	}
	return c
}
