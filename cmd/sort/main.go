package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/Cameron-Xie/golang-code/sort/bubble"
	"github.com/Cameron-Xie/golang-code/sort/merge"
)

func main() {
	s := getRandomIntSlice(100)

	fmt.Printf("merge sort: %q\n", getSliceStr(merge.Sort(s)))
	fmt.Printf("bubble sort: %q\n", getSliceStr(bubble.Sort(s)))
	
	fmt.Printf("original slice: %q\n", getSliceStr(s))
}

func getRandomIntSlice(n int) []int {
	rand.Seed(time.Now().UnixNano())
	s := make([]int, n)
	for i := range s {
		s[i] = rand.Intn(n)
	}

	return s
}

func getSliceStr(i []int) string {
	s := make([]string, 0)
	for _, n := range i {
		s = append(s, strconv.Itoa(n))
	}

	return strings.Join(s, ",")
}
