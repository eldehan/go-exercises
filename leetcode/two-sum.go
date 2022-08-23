// brute force - O(N^2)
// outer loop, iterate thru ints
// then, inner loop iterate thru other ints and add them to number of outer loop
// check if it hits target sum

// hash map - O(N)?
// add each # to hash map
// key is target-#, value is index
// at each step, you check whether current # is already in hash map
// if it is, you return the current index and the stored index in the hash map (which will be the match)
package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int)

	for i, v := range nums {
		if val, ok := hash[v]; ok {
			return []int{i, val}
		} else {
			key := target - v
			hash[key] = i
		}
	}
	return []int{0, 0}
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) // [1, 0]
}
