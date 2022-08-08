// exercise 11-1

// recursive func that accepts an array of strings, returns total # of
// chars across all the strings

// slice 1 element off array at a time and find the length of that element,
// keep calling func on the smaller slices until there's only 1 element left

package main

import "fmt"

func charCount(strSlice []string) int {
	// base case
	// when there is only 1 element left, take the length of that element
	if len(strSlice) == 1 {
		return len(strSlice[0])
	}
	// slice func from index 1 to the end, and call func again, adding its
	// return value to the length of the string at index 0
	return charCount(strSlice[1:]) + len(strSlice[0])
}

func main() {
	fmt.Println(charCount([]string{"ab", "c", "def", "ghi"})) // 9
	fmt.Println(charCount([]string{"only one string here"}))  // 20
	fmt.Println(charCount([]string{"1", "1", "1"}))           // 3
}
