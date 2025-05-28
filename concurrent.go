package main

import (
	"fmt"
	"sync"
)

func concurrent() {
	var threads sync.WaitGroup
	for i := 0; i < 32; i++ {
		threads.Add(1)
		go func(i int) {
			sendRPC(i)
			threads.Done()
		}(i)
	}
	threads.Wait()
}

func sendRPC(x int) {
	fmt.Println("sending request to server ", x)
}
