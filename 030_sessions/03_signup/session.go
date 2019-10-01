package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"net/http"
)

func getUser(w http.ResponseWriter, req *http.Request) user {
	// Get the Cookie from browser, based on the cookie return the user
	// if cookie absent create new cookie

	c, err := req.Cookie("session")
	if err != nil {
		sId, _ := uuid.NewV4()
		fmt.Println("Cookie not present", err)
		fmt.Println("new session id:", sId.String())
		c = &http.Cookie{
			Name:  "session",
			Value: sId.String(),
		}
		http.SetCookie(w, c)
	}

	var u user
	if userName, ok := dbSession[c.Value]; ok {
		u = dbUsers[userName]
	}
	return u
}

func isUserLoggedin(req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	userName := dbSession[c.Value]
	_, ok := dbUsers[userName]
	return ok
}
