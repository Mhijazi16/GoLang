// package main
//
// import (
// 	"fmt"
// )
//
// func sumNumber[T int | float32 | float64](slice []T) T {
// 	var sum T
// 	for i := range slice {
// 		sum += slice[i]
// 	}
// 	return sum
// }
//
// func main() {
//
// 	var one = []int{1, 2, 3, 4}
// 	var two = []float32{1.5, 2.8, 3.3, 4}
//
// 	fmt.Println("sum of integers: ", sumNumber(one))
// 	fmt.Println("sum of floats: ", sumNumber(two))
// }
