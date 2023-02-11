package main

func main() {
	
}

func run(nums []int) [][]int  {
	for i := 0; i < len(nums); i++ {
		list := []int{}
		//去掉 nums[i]

		run(list)
	}
}

