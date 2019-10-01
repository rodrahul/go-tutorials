package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	url2 "net/url"
	"strconv"
	"strings"
	"time"
)

type controls struct {
	DefaultMyCoursesFilter string `json:"defaultMyCoursesFilter"`
	HelpPage               string `json:"helppage"`
	HomePage               string `json:"homepage"`
	MinPasswordLength      string `json:"minPasswordLength"`
}

type loginPerson struct {
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

type userNamePassword struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func main() {
	domain := "https://jcp.helius.com"
	login := "/lms/authentication/login"
	urlString := domain + login

	g_domain := userNamePassword{
		Username: "g_domain",
		Password: "elearn",
	}

	bs, err := json.Marshal(g_domain)
	if err != nil {
		fmt.Println(err)
	}
	payload := strings.NewReader(string(bs))

	req, _ := http.NewRequest(http.MethodPost, urlString, payload)

	req.Header.Add("Accept", "application/json, text/plain, */*")
	req.Header.Add("Content-Type", "application/json")
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	//fmt.Println(string(body))

	var loginResponse loginPerson
	err = json.Unmarshal(body, &loginResponse)
	if err != nil {
		fmt.Println("error:", err)
	}

	//fmt.Println(loginResponse)

	fmt.Println("Token:", loginResponse.Token)

	//--- To get course Transcript-----
	//courseTranscript := "/lms/transcript/usertranscripts?personId=7404541"
	//urlString = domain + courseTranscript
	//req, _ = http.NewRequest(http.MethodGet, urlString, nil)
	//req.Header.Add("x-session-token", loginResponse.Token)
	//resp, err = client.Do(req)
	//if err != nil {
	//	panic(err)
	//}
	//defer resp.Body.Close()
	//body, _ = ioutil.ReadAll(resp.Body)
	//
	//fmt.Println(string(body))

	//------- To get users--------------
	userList := "/lms/user/list"
	urlString = domain + userList
	u, _ := url2.Parse(urlString)
	v := u.Query()
	v.Add("orgId", strconv.Itoa(loginResponse.OrgID))
	u.RawQuery = v.Encode()

	fmt.Println(u.String())
	req,_ = http.NewRequest(http.MethodGet, u.String(), nil)
	req, _ = http.NewRequest(http.MethodGet, urlString, nil)
	req.Header.Add("x-session-token", loginResponse.Token)
	req.Header.Add("Range", "records 0-1")
	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Printf("%+v\n",resp.Header)
	body, _ = ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

}
