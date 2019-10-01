package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"html/template"
	"io"
	"net/http"
	"os"
)

type user struct {
	UserName string
	First    string
	Last     string
}

var tpl *template.Template
var dbUsers = map[string]user{}      //[userId]user
var dbSessions = map[string]string{} //[sessionId] userId

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/favicon.ico", serveFavicon)

	http.ListenAndServe("localhost:8080", nil)
}

func serveFavicon(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("/Users/rrode/Hughes/System7/Themes/Iconfiles/camera.ico")
	if err != nil {
		http.Error(w, "favicon not found", 404)
		return
	}
	fmt.Println("Serving Favicon")
	defer f.Close()

	io.Copy(w, f)
}

func index(w http.ResponseWriter, req *http.Request) {
	// Get Cookie
	c, err := req.Cookie("session")
	if err != nil {
		sId, _ := uuid.NewV4()
		fmt.Println("Cookie was not set error is ", err.Error())
		c = &http.Cookie{
			Name:  "session",
			Value: sId.String(),
		}
		http.SetCookie(w, c)
	}
	// If user exists, get user
	var u user
	if un, ok := dbSessions[c.Value]; ok {
		u = dbUsers[un]
	}

	if req.Method == http.MethodPost {
		un := req.FormValue("username")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")

		u = user{un, f, l}
		dbSessions[c.Value] = u.UserName
		dbUsers[u.UserName] = u
	}
	tpl.ExecuteTemplate(w, "02_index.html", u)
	fmt.Println("Method:", req.Method)
	fmt.Println("User:", u)
}

func bar(w http.ResponseWriter, req *http.Request) {
	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	un, ok := dbSessions[c.Value]
	if !ok {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	u := dbUsers[un]
	tpl.ExecuteTemplate(w, "bar.html", u)
}
