// package main
//
// import (
// 	"fmt"
// )
//
// func main() {
// 	var slice = []int{1, 2, 3, 4}
// 	fmt.Printf("length: %v, capacity: %v\n", len(slice), cap(slice))
//
// 	slice = append(slice, 3)
// 	fmt.Printf("length: %v, capacity: %v\n", len(slice), cap(slice))
//
// 	var slice2 = []int{0, 0, 0, 0, 0}
// 	slice = append(slice, slice2...)
// 	fmt.Printf("length: %v, capacity: %v\n", len(slice), cap(slice))
//
// 	var slice3 []int = make([]int, 2, 9)
// 	fmt.Printf("length: %v, capacity: %v\n", len(slice3), cap(slice3))
// }
