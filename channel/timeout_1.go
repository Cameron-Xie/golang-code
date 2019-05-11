package main

import (
	"fmt"
	"time"
)

func main() {
	result := make(chan int, 1)

	//timeout := 3 // will timeout
	timeout := 10 // will find remote number

	go findMatchedNum(result)

	select {
	case res := <-result:
		fmt.Println("found it: ", res)
	case <-time.After(time.Duration(timeout) * time.Second):
		fmt.Println("timeout")
	}
}

func findMatchedNum(result chan int) {
	num := 0
	ticker := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-ticker.C:
			fmt.Printf("try num: %d\n", num)
			if remoteRequest(num) {
				result <- num
			}

			num++
		}
	}
}

func remoteRequest(num int) bool {
	secret := 5
	return num == secret
}
