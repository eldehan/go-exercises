package main

func maxSubArray(nums []int) int {
	return findMaxSubArray(nums, 0, len(nums)-1)
}

func findMaxSubArray(nums []int, left int, right int) int {
	// base case - left and right are the same value (can't split any further; one element only)
	if left == right {
		return nums[left]
	}

	// to find left and right halves, need a midpoint
	mid := left + ((right - left) / 2)

	// find left max
	// call findMaxSubArray on left half of array
	leftMax := findMaxSubArray(nums, left, mid)

	// find right max
	// call findMaxSubArray on right half of array
	rightMax := findMaxSubArray(nums, mid+1, right)

	// find crossover max - this is a contiguous series of #s between the two halves
	// pass in left and right maxes
	crossoverMax := findCrossoverMax(nums, left, right, mid)

	// find the max between left, right, and crossover
	return max(leftMax, rightMax, crossoverMax)
}

func findCrossoverMax(nums []int, left int, right int, mid int) int {
	// need to define left and right maxes which start out extremely negative
	// because initial comparisons have to be bigger
	leftMax := -99999
	rightMax := -99999

	// when building up the left and right maxes here, we build outwards from mid, heading left (for left)
	// and right (for right)
	// this is to build 2 contiguous results we can add together to get crossover

	// if current running sum is > than our saved leftMax, it becomes the new leftMax
	sum := 0
	for i := mid; i >= left; i-- {
		sum += nums[i]
		leftMax = max(leftMax, sum)
	}

	// do the same for rightMax
	sum = 0
	for i := mid + 1; i <= right; i++ {
		sum += nums[i]
		rightMax = max(rightMax, sum)
	}

	return leftMax + rightMax
}

func max(values ...int) int {
	// again, set initial value to very negative
	maxValue := -99999

	// if any value is bigger than the current max, it becomes new max
	for i := 0; i < len(values); i++ {
		if values[i] > maxValue {
			maxValue = values[i]
		}
	}

	return maxValue
}
