package huawei

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {

}

//标准输入
func input() {
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	input := string(data)
	fmt.Println(input)
}

//排序
func sortDesc() {
	a := []int{1,4,5,7}
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})
}