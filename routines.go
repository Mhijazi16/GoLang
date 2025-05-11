package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var database = []string{"one", "two", "three", "four"}
var tasks = sync.WaitGroup{}

func getElement(id int) {
	var delay = rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Printf("retrived %v from the database\n", database[id])
	tasks.Done()
}

func main() {
	for index := range database {
		tasks.Add(1)
		go getElement(index)
	}

	tasks.Wait()
	fmt.Println("all elements retrieved!")
}
