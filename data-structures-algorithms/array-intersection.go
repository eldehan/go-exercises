// exercise 8-1

package main

import "fmt"

// constructs lookup map to streamline efficiency of algorithm to O(N)
func findIntersection(arr1, arr2 []int) []int {
	// initialize map
	lookup := make(map[int]bool)

	// for each value in first slice, store that value in the map
	for _, v := range arr1 {
		lookup[v] = true
	}

	// initialize result slice
	resultArr := []int{}

	// for each value in the second slice, check to see if it's in the map
	// then add it to the result slice
	for _, v := range arr2 {
		if lookup[v] {
			resultArr = append(resultArr, v)
		}
	}

	return resultArr
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{0, 2, 4, 6, 8}

	fmt.Println(findIntersection(arr1, arr2)) // [2 4]
}
