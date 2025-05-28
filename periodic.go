package main

import (
	"fmt"
	"sync"
	"time"
)

func routine() {
	fmt.Println("starting routine!")
	time.Sleep(1 * time.Second)

	var threads sync.WaitGroup
	threads.Add(1)
	go periodic(&threads)

	threads.Wait()
}

func periodic(threads *sync.WaitGroup) {

	fmt.Println("started periodic cron job")
	for range 5 {
		time.Sleep(1 * time.Second)
		fmt.Println("finished some task")
	}

	threads.Done()
}

