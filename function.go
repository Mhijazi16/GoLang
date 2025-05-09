// package main
//
// import (
// 	"errors"
// 	"fmt"
// )
//
// func intDivision(numerator int, denomirator int) (int, int, error) {
// 	if denomirator == 0 {
// 		var err = errors.New("division by zero not allowed")
// 		return 0, 0, err
// 	}
//
// 	return numerator / denomirator, numerator % denomirator, nil
// }
//
// func main() {
//
// 	var arr = [2][2]int{{5, 2}, {0, 0}}
//
// 	for i := range 2 {
// 		var first, second, err = intDivision(arr[i][0], arr[i][1])
//
// 		if err != nil {
// 			fmt.Printf("failure! resulsts %v %v\n", first, second)
// 			fmt.Println(err.Error())
// 		} else {
// 			fmt.Printf("success resulsts %v %v", first, second)
// 		}
// 	}
//
// }
