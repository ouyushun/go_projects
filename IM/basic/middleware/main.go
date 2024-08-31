package main

import (
	"fmt"
	"net/http"
)

func Test1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test1"))
}
func Test2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("test2"))
}


func main() {
	http.HandleFunc("/", Test1)
	//http.ListenAndServe("127.0.0.1:9999", nil, )


	err := http.ListenAndServe("127.0.0.1:12345", http.FileServer(http.Dir(".")))
	if err != nil {
		fmt.Println(err)
		return
	}
}
