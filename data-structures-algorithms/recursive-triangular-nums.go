// exercise 11-3

// triangular numbers are a pattern that begins:
// 1, 3, 6, 10, 15, 21, etc
// the Nth # is N + the previous #

package main

import "fmt"

func triangularNums(n int) int {
	if n == 1 {
		return 1
	}
	return n + triangularNums(n-1)
}

func main() {
	fmt.Println(triangularNums(7)) // 28
	fmt.Println(triangularNums(6)) // 21
	fmt.Println(triangularNums(5)) // 15
}
