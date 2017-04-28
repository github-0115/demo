package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Person struct {
	Name string `json:"name"`
	Age  int32  `json:"age"`
}

func (p *Person) String() string {
	return "name: " + p.Name + " ，age: " + strconv.Itoa(int(p.Age))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	var res string = "hello :\n"
	for k, vs := range r.URL.Query() {
		for _, v := range vs {
			res += "\t" + k + ":" + v + " !\n"
		}
	}
	io.WriteString(w, res)
}

func languageHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	checkErr(err)

	var p *Person
	err = json.Unmarshal(body, &p)
	checkErr(err)
	fmt.Println(p.String())

	io.WriteString(w, string(body))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/language", languageHandler)

	http.ListenAndServe(":8000", mux)
}

//error
func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
