package main

import (
	"fmt"
	"strings"

	"github.com/Cameron-Xie/golang-code/puzzle/jsonparser"
	"github.com/Cameron-Xie/golang-code/puzzle/numsteps"
)

func main() {
	/*
		stairs: 4
		steps: 1 or 2

		ways:
		1111
		112
		211
		121
		22
	*/
	stairs, steps := 4, []int{1, 2}
	ways := numsteps.Calculate(stairs, steps)
	fmt.Printf("steps %v, total %d ways to reach %d stairs\n", steps, ways, stairs)

	/*
		input: {"foo":"bar\\"}{"foo2 {}":"bar2"}
		output:
		{"foo":"bar\\"}
		{"foo2 {}":"bar2"}
	*/
	jsonSeq := `{"foo":"bar\\"}{"foo2 {}":"bar2"}`
	fmt.Printf("JSON-seq parsing output %v\n", strings.Join(jsonparser.ParseJSONSeq(jsonSeq), ", "))
}
