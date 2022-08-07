// exercise 8-3

// func that takes a string which has all the letters of the alphabet except one
// return the missing letter
// should have a time complexity of O(N)

// initialize map of alphabet with all letters seen
// convert string to slice of strings

package main

import (
	"fmt"
	"strings"
)

func missingLetter(input string) string {
	alphabet := map[string]bool{"a": true, "b": true, "c": true, "d": true, "e": true, "f": true, "g": true, "h": true, "i": true, "j": true, "k": true, "l": true, "m": true, "n": true, "o": true, "p": true, "q": true, "r": true, "s": true, "t": true, "u": true, "v": true, "w": true, "x": true, "y": true, "z": true}

	lettersSlice := strings.Split(input, "")
	for _, v := range lettersSlice {
		delete(alphabet, v)
	}

	// sort of messy; if multiple, would technically would just return the first one encountered in the range
	// but, can infer from prompt that there will only be 1 missing letter
	for key := range alphabet {
		return key
	}

	return "no missing letters"
}

func main() {
	fmt.Println(missingLetter("the quick brown box jumps over a lazy dog"))
}
