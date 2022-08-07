package main

import "fmt"

// n is starting point
// just loop and count down and append it to a slice

func ReverseSeq(n int) []int {
	results := []int{}
	for countDown := n; countDown > 0; countDown-- {
		results = append(results, countDown)
	}
	return results
}

func main() {
	fmt.Println(ReverseSeq(10))
}
