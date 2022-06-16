package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	res := `{
		"data":[
			{
				"a":"b"
			}
		]
	}
	`

	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(res), &m)
	if err != nil {
		//panic(err)
		fmt.Println(err)
		return
	}
	fmt.Println([]byte(res))
	fmt.Printf("%T", m["data"].([]interface{})[0])
}
