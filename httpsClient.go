package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "https://postman-echo.com/post"
	fmt.Println("URL:>", url)

	var jsonStr = []byte(`{"foo":"bar"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println()
	fmt.Println("response Headers:", resp.Header)
	fmt.Println()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	url2 := "https://jcp.helius.com/login/authenticate"
	fmt.Println("URL:>", url2)
}
