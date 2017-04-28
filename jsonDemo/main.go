package main

import (
	"encoding/json"
	"fmt"
)

/*
* CompanyName : 公司名称
* StaffCount : 职员数量
* StaffNames : 职员名字
 */
type Company struct {
	CompanyName string   `json:"companyName"`
	StaffCount  int32    `json:"staffCount"`
	StaffNames  []string `json:"staffNames"`
}

func main() {

	marshal()

	unmarshal()
}

//对象序列化为json字符串
func marshal() {
	company := &Company{
		CompanyName: "golang",
		StaffCount:  2,
		StaffNames:  []string{"go1", "go2"},
	}

	b, err := json.Marshal(company)
	checkErr(err)
	if b != nil {
		fmt.Println(string(b))
	}
}

//json字符串反序列化为对象
func unmarshal() {
	var (
		company *Company
		jsonStr string
	)
	jsonStr = `{"companyName":"golang","staffCount":2,"staffNames":["go1","go2"]}`
	err := json.Unmarshal([]byte(jsonStr), &company)
	checkErr(err)
	if err == nil {
		fmt.Println("CompanyName:", company.CompanyName)
	}
}

//error
func checkErr(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
}
