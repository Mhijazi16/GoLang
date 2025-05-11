package main

import (
	"fmt"
)

func sum(nums ...int) int {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	return sum
}

func main() {
	fmt.Println(sum(1))
	fmt.Println(sum(1, 2))
	fmt.Println(sum(1, 2, 3))

	var arr = []int{1, 2, 3, 4}
	fmt.Println(sum(arr...))
}
