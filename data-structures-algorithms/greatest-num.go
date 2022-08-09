// write 3 different implementations of a function that
// finds the greatest # in an array of ints
// efficiencies should be O(N^2), O(N log N), and O(N)

package main

import (
	"fmt"
	"sort"
)

func greatestNSquared(nums []int) int {
	for i := 0; i < len(nums); i += 1 {
		iIsGreatest := true
		for j := i + 1; j < len(nums); j += 1 {
			if nums[j] > nums[i] {
				iIsGreatest = false
			}
		}

		if iIsGreatest {
			return nums[i]
		}
	}

	return 0
}

func greatestNLogN(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)-1]
}

func greatestN(nums []int) int {
	greatest := nums[0]
	for _, v := range nums {
		if v > greatest {
			greatest = v
		}
	}

	return greatest
}

func main() {
	fmt.Println(greatestNSquared([]int{2, 3, 7, 9}))
	fmt.Println(greatestNLogN([]int{2, 3, 7, 9}))
	fmt.Println(greatestN([]int{2, 3, 7, 9}))
}
