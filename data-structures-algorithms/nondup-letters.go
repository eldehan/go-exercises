// exercise 8-4

package main

import (
	"fmt"
	"strings"
)

func findNonDups(str string) string {
	letterTracker := map[string]int{}
	stringSlice := strings.Split(str, "")

	// checks whether value exists in map yet or not; if it does, increment value by 1; if not, add it as a key with a value of 1
	for _, letter := range stringSlice {
		if _, ok := letterTracker[letter]; ok {
			letterTracker[letter] += 1
		} else {
			letterTracker[letter] = 1
		}
	}

	// iterate through map and return the first letter encountered with a value of 1
	for key, val := range letterTracker {
		if val == 1 {
			return key
		}
	}

	return "no non-duplicates"
}

func main() {
	fmt.Println(findNonDups("minimum"))
}
