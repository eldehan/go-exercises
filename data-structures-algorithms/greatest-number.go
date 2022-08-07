// Exercise 4-4

package main

import "fmt"

// initial, given function
// efficiency of O(N^2)
func greatestNumber(arr []int) int {
	// start off assuming the indexed # is the greatest
	// loop through the numbers
	for _, i := range arr {
		isIvalTheGreatest := true

		// loop through the rest of the #s; if any are greater, change greatest variable to false
		for _, j := range arr {
			if j > i {
				isIvalTheGreatest = false
			}
		}

		// if the i value is still the greatest after comparing to the rest of the numbers, return that number
		if isIvalTheGreatest {
			return i
		}
	}

	return 0
}

// Revised function with time complexity of O(N)
func greatestNumberRevised(arr []int) int {
	// set greatest seen num variable to the first array item
	greatestSeenNum := arr[0]

	for _, v := range arr {
		// Loop through the array once; if any number is greater than the currently greatest seen #, it takes its place
		if v > greatestSeenNum {
			greatestSeenNum = v
		}
	}

	return greatestSeenNum
}

func main() {
	fmt.Println(greatestNumber([]int{1, 2, 3, 4, 5}))        // 5
	fmt.Println(greatestNumber([]int{123, 24, 3, 42, 5}))    // 123
	fmt.Println(greatestNumber([]int{1, 72, 93, 2314, 555})) // 2314

	fmt.Println(greatestNumberRevised([]int{1, 2, 3, 4, 5}))        // 5
	fmt.Println(greatestNumberRevised([]int{123, 24, 3, 42, 5}))    // 123
	fmt.Println(greatestNumberRevised([]int{1, 72, 93, 2314, 555})) // 2314
}
