package main

import (
	_ "debug/dwarf"
	"fmt"
	"net/http"
	"net/http/httputil"
)

func main() {
	resp, err := http.Get("http://www.imooc.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Body)
	fmt.Println(err)
	s, err := httputil.DumpResponse(resp, true)
	fmt.Printf("%s", s)


}
