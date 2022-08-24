// binary search program to find a number in an array of sorted numbers from scratch
// so:
// numbers already sorted
// presumably, return index if found, or 0 if not found

// find length of array, then find halfway point
// compare halfway item vs sought item
// if item == sought item, you're done;
// if item > target, set end to midpoint-1
// if item < target, set start to midpoint+1

package main

import "fmt"

func binarySearch(numbers []int, target int) int {
	// start off setting 'found' index to -1, indicating it's not found yet
	index := -1

	// start and end assigned to 0 index and last index respectively
	start, end := 0, len(numbers)-1

	// loop to run as long as the start is not greater than the end
	for start <= end {
		midpoint := (start + end) / 2
		// if target is found, set the 'found index' to the result and break the loop
		if numbers[midpoint] == target {
			index = midpoint
			break
		} else if numbers[midpoint] < target {
			start = midpoint + 1
		} else {
			end = midpoint - 1
		}
	}
	return index
}

func main() {
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5}, 5))
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5}, 6))
	fmt.Println(binarySearch([]int{}, 6))
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5, 6}, 3))
}
