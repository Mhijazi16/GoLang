// package main
//
// // we should have two routines one checks for checkn prices
// // the second sends a message when price is lower than threshold
//
// // first function checks price and appends website value to channel if price < threshold
// // second function reads the value of chaneel and send message when value supplied
//
// import (
// 	"fmt"
// 	"math/rand"
// 	"time"
// )
//
// func checkCpuPrices(website string, channel chan string) {
// 	for {
// 		time.Sleep(time.Second * 2)
// 		if rand.Intn(100) >= 20 {
// 			channel <- website
// 			break
// 		}
// 	}
// }
//
// func sendMessage(cpuChannel chan string) {
// 	fmt.Println("price found for cpu in:", <-cpuChannel)
// }
//
// func main() {
// 	cpuChannel := make(chan string)
//
// 	var websites = []string{"nvidiaCPU", "intelCPU", "AMDCPU"}
// 	for index := range websites {
// 		go checkCpuPrices(websites[index], cpuChannel)
// 		sendMessage(cpuChannel)
// 	}
// }
