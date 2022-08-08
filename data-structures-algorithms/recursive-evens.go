// exercise 11-2

// recursive func that accepts an array of #s and returns a new array containing
// only the even numbers

package main

import "fmt"

func recursiveEvens(numArr []int) []int {
	if len(numArr) == 0 {
		return []int{}
	}
	if numArr[0]%2 == 0 {
		return append([]int{numArr[0]}, recursiveEvens(numArr[1:])...)
	} else {
		return recursiveEvens(numArr[1:])
	}
}

func main() {
	fmt.Println(recursiveEvens([]int{1, 2, 3, 4, 5})) // [2 4]
}
