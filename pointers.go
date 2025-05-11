package main

import (
	"fmt"
)

func printNumbers(arr *[]int32) {
	fmt.Printf("memory address %p\n", arr)
	for index, value := range *arr {
		fmt.Printf("[%v] %v\n", index, value)
	}
}

func pointers() {
	var arr = []int32{1, 3, 4, 5, 6}
	fmt.Printf("memory address %p\n", &arr)
	printNumbers(&arr)
}
