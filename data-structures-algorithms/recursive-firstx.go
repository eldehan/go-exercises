// exercise 11-4
// return index of first x encountered
// can assume there will always be an x in the string
// (no guard clause needed)

package main

import "fmt"

func firstX(str string, index int) int {
	if str[index] == 'x' {
		return index
	} else {
		return firstX(str, index+1)
	}
}

func main() {
	fmt.Println(firstX("abcdefghijklmnopqrstuvwxyz", 0))
}
