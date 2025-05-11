package main

import (
	"fmt"
	"time"
)

func timeLoop(slice []int, n int) {
	var start = time.Now()

	for len(slice) < n {
		slice = append(slice, 0)
	}
	fmt.Printf("time taken: %v\n", time.Since(start))
}

func time_go() {

	n := 10000000
	var one = make([]int, 1, n)
	var two = []int{}

	timeLoop(one, n)
	timeLoop(two, n)

}
