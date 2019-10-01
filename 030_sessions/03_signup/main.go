package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"html/template"
	"io"
	"net/http"
	"os"
)

type user struct {
	UserName string
	Password string
	First    string
	Last     string
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/signup", signUp)
	http.HandleFunc("/favicon.ico", favicon)

	http.ListenAndServe(":8080", nil)
}

var tpl *template.Template
var dbUsers = map[string]user{}     // userId -> username
var dbSession = map[string]string{} // session -> userId

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func foo(w http.ResponseWriter, req *http.Request) {
	// Get user
	u := getUser(w, req)
	tpl.ExecuteTemplate(w, "index.html", u)
	fmt.Println("-----At root-----------------")
	fmt.Println("dbSession:", dbSession)
	fmt.Println("dbUsers:", dbUsers)
	fmt.Println("-----------------------------")

}

func bar(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	if !isUserLoggedin(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}
	tpl.ExecuteTemplate(w, "bar.html", u)
}

func signUp(w http.ResponseWriter, req *http.Request) {
	if isUserLoggedin(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	// process form submission
	if req.Method == http.MethodPost {

		// get form values
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")

		// username taken?
		if _, ok := dbUsers[un]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		// create session
		sID, _ := uuid.NewV4()
		c := &http.Cookie{
			Name:  "session",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
		dbSession[c.Value] = un

		// store user in dbUsers
		u := user{un, p, f, l}
		dbUsers[un] = u

		// redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "signup.html", nil)
}

func favicon(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("/Users/rrode/Hughes/System7/Themes/Iconfiles/camera.ico")
	if err != nil {
		http.Error(w, "cannot serve favicon", http.StatusNotFound)
	}
	io.Copy(w, f)
}
