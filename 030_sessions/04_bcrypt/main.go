package main

import (
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"net/http"
)

type user struct {
	UserName string
	password []byte
	First    string
	Last     string
}

var tpl *template.Template

var dbSessions = map[string]string{} // sessionId -> userName
var dbUsers = map[string]user{}      //userName -> user

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/signup", signUp)
	http.HandleFunc("/bar", bar)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	http.ListenAndServe("localhost:8080", nil)
}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func index(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	tpl.ExecuteTemplate(w, "index.html", u)

}

func bar(w http.ResponseWriter, req *http.Request) {
	u := getUser(w, req)
	if !isUserLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	tpl.ExecuteTemplate(w, "bar.html", u)
}

func signUp(w http.ResponseWriter, req *http.Request) {
	if isUserLoggedIn(req) {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	var u user
	if req.Method == http.MethodPost {
		// Get form values
		un := req.FormValue("username")
		p := req.FormValue("password")
		f := req.FormValue("firstname")
		l := req.FormValue("lastname")

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
		dbSessions[c.Value] = un

		bs, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		u = user{un, bs, f, l}
		dbUsers[un] = u

		// redirect
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return

	}
	tpl.ExecuteTemplate(w, "signup.html", nil)
}
