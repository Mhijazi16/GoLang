package main

import (
	"fmt"
)

type engine interface {
	milesLeft(float32) float32
}

func canMakeIt(eg engine, distance float32) float32 {
	return eg.milesLeft(distance) * 10
}

type gasEngine struct {
	speed   float32
	gallons int32
}

func (self gasEngine) milesLeft(distance float32) float32 {
	return self.speed / distance
}

type electricalEngine struct {
	speed       float32
	accelration float32
}

func (self electricalEngine) milesLeft(distance float32) float32 {
	return (self.speed * self.accelration) / distance
}

func structs() {

	var one = gasEngine{speed: 20}
	var two = electricalEngine{speed: 20, accelration: 2}

	fmt.Println(canMakeIt(one, 100))
	fmt.Println(canMakeIt(two, 100))
}
