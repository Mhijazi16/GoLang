package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now().Weekday()

	switch now {
	case time.Sunday, time.Saturday:
		fmt.Println("weekend")
	default:
		fmt.Println("work day")
	}

	FindType := func(abc interface{}) {
		switch abc.(type) {
		case bool:
			fmt.Println("I am a bool")
		case int:
			fmt.Println("I am a integer")
		default:
			fmt.Println("I am a string")
		}
	}

	FindType(true)
	FindType(2)
	FindType("hello")
}
