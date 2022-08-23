// [-1,0,1,2,-1,-4] sort

// set up fixed/left/right (fixed starts at 0; left at fixed+1; right at len-1)
// [-4,-1,-1,0,1,2]
//   f  l        r

// add together; if too small, increment left; if too big, decrement right;
// add solutions as you find them
// skip over duplicate values from both ends to avoid duplicate triplets

// iterate through by incrementing fixed

package main

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	solutions := [][]int{}

	for fixedIndex, val := range nums {
		// if the value at fixedIndex is the same as the prev value used as fixedIndex, continue; avoid the dup
		if fixedIndex > 0 && nums[fixedIndex] == nums[fixedIndex-1] {
			continue
		}

		left := fixedIndex + 1
		right := len(nums) - 1

		for left < right {
			sum := val + nums[left] + nums[right]

			if sum == 0 {
				solutions = append(solutions, []int{val, nums[left], nums[right]})
				right -= 1

				// if the new value of right is the same as the prev value of right, keep decrementing right until it is not
				// avoid duplicates from the right
				for left < right && nums[right] == nums[right+1] {
					right -= 1
				}
			} else if val+nums[left]+nums[right] < 0 {
				left += 1
			} else {
				right -= 1
			}
		}
	}

	return solutions
}
