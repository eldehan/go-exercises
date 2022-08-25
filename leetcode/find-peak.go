// binary search
// target in this case will be if the mid is greater than both mid-1 and mid+1, otherwise...
// keep running binary searches!

package main

func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1

	for right > left {
		mid := (left + right) / 2
		// if mid is greater than mid + 1, that means a peak is to the left (or is mid)
		if nums[mid] > nums[mid+1] {
			right = mid
			// otherwise, if mid is LESS THAN mid + 1, that means the peak is to the right
			// so, move left boundary up to keep searching to the right
		} else {
			left = mid + 1
		}
		// eventually, you'll have nums[left] > nums[right], and left and right will be next to each other
		// this is your peak
	}
	return left
}
