package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"net/http"
)

func isUserLoggedIn(req *http.Request) bool {
	c, err := req.Cookie("session")
	if err != nil {
		return false
	}
	un := dbSessions[c.Value]
	_, ok := dbUsers[un]
	return ok
}

func getUser(w http.ResponseWriter, req *http.Request) user {
	var u user
	c, err := req.Cookie("session")
	if err != nil {
		sId, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sId.String(),
		}
		http.SetCookie(w, c)
		fmt.Println("New Session, cookie is not present, err is:", err)
		fmt.Println("Cookie: ", c.Value)
	}

	// Check if user is present for the cookie

	if userName, ok := dbSessions[c.Value]; ok {
		u = dbUsers[userName]
	}

	fmt.Println("dbSession: ", dbSessions)
	fmt.Println("dbUsers:")
	for _, v := range dbUsers {
		fmt.Println(v.UserName, v.First, v.Last)
	}

	return u
}
