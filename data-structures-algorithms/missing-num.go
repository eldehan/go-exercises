// exercise 13-2

// given a func that, given an array of ints, finds the
// 'missing' number. i.e., the numbers in the array are
// in sequence from 0 and one that SHOULD be there is not
// initial implementation was time complexity of O(N^2)
// rewrite so it is O(N log N)

package main

import (
	"fmt"
	"sort"
)

func missingNum(nums []int) int {
	sort.Ints(nums)
	for count := 0; count < len(nums); count += 1 {
		if nums[count] != count {
			return count
		}
	}

	return 0
}

func main() {
	fmt.Println(missingNum([]int{9, 3, 2, 5, 6, 7, 1, 0, 4})) // 8
}
