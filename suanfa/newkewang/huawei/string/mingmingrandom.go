package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	total, _ := strconv.Atoi(string(data))
	inputMap := make(map[int]int)
	for i := 0; i < total ; i++ {
		data, _, _ = reader.ReadLine()
		v, _ := strconv.Atoi(string(data))
		inputMap[v] = v
	}
	dataSort := make([]int, 0)
	for _, v := range inputMap {
		dataSort = append(dataSort, v)
	}

	sort.Slice(dataSort, func(i, j int) bool {
		return dataSort[i] > dataSort[j]
	})
	for _, vv := range dataSort {
		fmt.Println(vv)
	}
}
