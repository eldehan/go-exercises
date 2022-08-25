// rotate array: has 2 halves that are ordered
// if last value is less than the middle, then go right otherwise go left
// if value on left and right are both greater, then you have the minimum value

// binary searchhhhhh
// take middle
// if mid > mid-1 and mid < mid+1, you have the min
// last value of array less than middle?
// yes - left = mid+1
// no - right = mid
// this will eventually slot the min into the right element

package main

func findMin(nums []int) int {
	last := len(nums) - 1
	left, right := 0, last

	for right > left {
		mid := (left + right) / 2

		if nums[mid] > nums[last] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[right]
}
