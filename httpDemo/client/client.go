package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

var (
	helloUrl    string = "http://localhost:8000/hello?name=goalng&age=9"
	languageUrl string = "http://localhost:8000/language"
)

func main() {

	client := GetClient()

	DoGet(client, helloUrl)
	fmt.Println()

	parmar := `{"name":"golang","age":9}`
	DoPost(client, languageUrl, parmar)
}

func GetClient() *http.Client {
	client := &http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).Dial,
			TLSHandshakeTimeout:   10 * time.Second,
			ResponseHeaderTimeout: 10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		}}

	return client
}

func DoGet(client *http.Client, url string) string {
	resp, err := client.Get(url)
	checkErr(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)
	fmt.Printf("%s", string(body))
	return string(body)
}

func DoPost(client *http.Client, url string, parmar string) string {

	resp, err := client.Post(url,
		"application/json; charset=utf-8",
		strings.NewReader(parmar))
	checkErr(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err)
	fmt.Println(string(body))
	return string(body)
}

//error
func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
