package main

import (
	"fmt"
	"io/ioutil"
	http "net/http"
	"regexp"
)

func main() {
	response, err := http.Get("http://localhost:8080/mock/www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		fmt.Println("error, status:", response.StatusCode)
	}
	all, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%T", all)
	printCityList(all)
}

func printCityList(contents []byte) {
	reg := regexp.MustCompile(`<a href="(http://localhost:8080/mock/www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)
	//fmt.Println(reg)
	matchsRes := reg.FindAllSubmatch(contents, -1)
	//fmt.Printf("%s", matchsRes)

	for _, m := range matchsRes {
		fmt.Printf("%s %s", m[1], m[2])
		fmt.Println()
	}

	println(len(matchsRes))
}
