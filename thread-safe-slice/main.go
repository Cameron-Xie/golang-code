package main

import (
	"math/rand"
	"sync"
	"time"
)

const (
	minInt = 1
	maxInt = 10
)

func main() {

}

func getRandomInt(min, max int) int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(max-min+1) + min
}

func LockAppend(s []int, num int) []int {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	wg.Add(num)
	for i := 1; i <= num; i++ {
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Duration(getRandomInt(minInt, maxInt)*100) * time.Millisecond)
			mu.Lock()
			defer mu.Unlock()
			s = append(s, i)
		}(i)
	}

	wg.Wait()

	return s
}

func ChannelAppend(s []int, num int) []int {
	c := make(chan int)

	for i := 1; i <= num; i++ {
		go func(v int) {
			time.Sleep(time.Duration(getRandomInt(minInt, maxInt)*100) * time.Millisecond)
			c <- v
		}(i)
	}

	for i := range c {
		s = append(s, i)

		if len(s) == num {
			close(c)
		}
	}

	return s
}
