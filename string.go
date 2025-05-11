package main

import (
	"fmt"
	"strings"
)

func main() {
	var mystring = "H😾L"

	for index, value := range mystring {
		fmt.Printf("[%v] value: %v, Type: %T\n", index, value, value)
	}

	fmt.Println("=======")

	var runes = []rune(mystring)
	for index, value := range runes {
		fmt.Printf("[%v] value: %v, Type: %T\n", index, value, value)
	}

	fmt.Println("=======")

	var letters = []string{"H", "😾", "L"}
	var builder strings.Builder
	for i := range letters {
		builder.WriteString(letters[i])
	}

	var final = builder.String()
	fmt.Println(final)

}
