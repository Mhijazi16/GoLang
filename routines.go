package main

import (
	"fmt"
	"sync"
	"time"
)

var database = []string{"one", "two", "three", "four"}
var results = []string{}
var tasks = sync.WaitGroup{}
var mutex = sync.Mutex{}

func getElement(id int) {
	var delay = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Printf("retrived %v from the database\n", database[id])
	mutex.Lock()
	results = append(results, database[id])
	mutex.Unlock()

	tasks.Done()
}

func main() {
	for index := range database {
		tasks.Add(1)
		go getElement(index)
	}

	tasks.Wait()

	fmt.Println("all elements retrieved!")

	for index := range results {
		fmt.Println(results[index])
	}
}
