package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type User struct {
	Name  string `json:"user_name"`
	Age   int    `json:"age"`
	sex   string
	Work1 *Work1
	Work2 json.RawMessage
	Work3 Work3
	Work4 interface{} `json:"work4"`
	Work5 interface{}
}

type Work1 struct {
	Name   string `json:"work1_name"`
	Salary float32
}

type Work2 struct {
	Name   string `json:"work2_name"`
	Salary float32
}

type Work3 struct {
	Name   string `json:"work3_name"`
	Salary float32
}

type Work4 struct {
	Name    string `json:"work4_name"`
	Salary  float32
	Address string `json:"work4_add"`
}

func main() {
	//json字符中的"引号，需用\进行转义，否则编译出错
	data := "{\"user_name\":\"ares\",\"sex\":\"男\",\"age\":18,\"Work1\":{\"work1_name\":\"god1\",\"Salary\":100},\"Work2\":{\"work2_name\":\"god2\",\"Salary\":200},\"Work3\":{\"work3_name\":\"god3\",\"Salary\":300},\"work4\":{\"work4_name\":\"god4\",\"Salary\":400,\"work4_add\":\"cbd\"}}"
	str := []byte(data)
	u1 := User{}
	// Unmarshal的第一个参数是json字符串，第二个参数是接受json解析的数据结构.第二个参数必须是指针，否则无法接收解析的数据，
	err := json.Unmarshal(str, &u1)
	if err != nil {
		fmt.Println("Unmarshal err,", err)
	}
	// {ares 18  0xc0000a41c8 0xc0000a41e0 {god3 300} map[Salary:400 work4_add:cbd work4_name:god4]}  Work2 为*Work2类型
	//  Work2 为json.RawMessage类型  {ares 18  0xc0000a4198 [123 34 119 111 114 107 50 95 110 97 109 101 34 58 34 103 111 100 50 34 44 34 83 97 108 97 114 121 34 58 50 48 48 125] {god3 300} map[Salary:400 work4_add:cbd work4_name:god4] <nil>}

	fmt.Println(u1)
	// 查看类型
	nameType := reflect.TypeOf(u1.Name)
	ageType := reflect.TypeOf(u1.Age)
	sexType := reflect.TypeOf(u1.sex)
	work1Type := reflect.TypeOf(u1.Work1)
	work2Type := reflect.TypeOf(u1.Work2)
	work3Type := reflect.TypeOf(u1.Work3)
	work4Type := reflect.TypeOf(u1.Work4)
	work5Type := reflect.TypeOf(u1.Work5)
	fmt.Println(nameType)  // string
	fmt.Println(ageType)   // int
	fmt.Println(sexType)   // string
	fmt.Println(work1Type) // *main.Work1
	fmt.Println(work2Type) // json.RawMessage
	fmt.Println(work3Type) // main.Work3
	fmt.Println(work4Type) // map[string]interface {}
	fmt.Println(work5Type) // <nil>
}