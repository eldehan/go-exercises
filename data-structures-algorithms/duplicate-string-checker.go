// exercise 8-2

// func should accept an array of strings and return the first duplicate value that it finds
// i.e. in an array of ["a", "b", "c", "d", "c", "e", "f"], it should return "c"
// can assume there is one pair of duplicates within the array
// efficiency should be O(N)

// create map
// iterate through array and check to see if the string is already in the map
// if not, store the seen strings in the map as keys
// if yes, return that string immediately

package main

import "fmt"

func duplicateChecker(stringSlice []string) string {
	seenMap := make(map[string]bool)

	for _, v := range stringSlice {
		if !seenMap[v] {
			seenMap[v] = true
		} else {
			return v
		}
	}

	return "no duplicates"
}

func main() {
	testSlice := []string{"a", "b", "c", "d", "c", "e", "f"}
	testSlice2 := []string{"apple", "pear", "no duplicates here"}

	fmt.Println(duplicateChecker(testSlice))  // c
	fmt.Println(duplicateChecker(testSlice2)) // no duplicates
}
