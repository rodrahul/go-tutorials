package login

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"gopdf/mtTesting/api"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type controls struct {
	DefaultMyCoursesFilter string `json:"defaultMyCoursesFilter"`
	HelpPage               string `json:"helppage"`
	HomePage               string `json:"homepage"`
	MinPasswordLength      string `json:"minPasswordLength"`
}

type loginApiResponse struct {
	PersonID               int       `json:"personId"`
	DomainID               int       `json:"domainId"`
	OrgID                  int       `json:"orgId"`
	DefaultMyCoursesFilter string    `json:"defaultMyCoursesFilter"`
	Rights                 []string  `json:"rights"`
	Controls               controls  `json:"controls"`
	FirstPage              string    `json:"firstPage"`
	FirstName              string    `json:"firstName"`
	LastName               string    `json:"lastName"`
	MiddleInitial          string    `json:"middleInitial"`
	Image                  string    `json:"image"`
	Language               string    `json:"language"`
	VideoPlayerLicense     string    `json:"videoPlayerLicense"`
	Token                  string    `json:"token"`
	RemoteHost             string    `json:"remoteHost"`
	ExpiryTime             time.Time `json:"expiryTime"`
	RemoteAddress          string    `json:"remoteAddress"`
	OrgFolderID            int       `json:"orgFolderId"`
}

type Person struct {
	LoginResponse loginApiResponse
	HTTPClient    *http.Client
}

func Login(userName, password string) Person {
	unP := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{userName, password}

	bs, _ := json.Marshal(unP)
	payload := strings.NewReader(string(bs))

	req, _ := http.NewRequest(http.MethodPost, api.LoginAPI, payload)
	req.Close = true
	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("Content-Type", "application/json")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		DisableKeepAlives:true,
	}

	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	var loginResponse loginApiResponse
	err = json.Unmarshal(body, &loginResponse)
	if err != nil {
		fmt.Println("error:", err)
	}
	personDetails := Person{loginResponse, client}

	return personDetails
}
