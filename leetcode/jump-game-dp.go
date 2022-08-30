// for each new dp spot, compare whether the previous jump
// or the current jump we're using gives us more remaining spaces
// to travel (-1, to account for moving forward 1 spot)

// if at any point, after we calculate our new jump capacity,
// dp[i] is less than 0, return false (cannot move on)

// at the end, return whether we made it to the target

package main

func canJump(nums []int) bool {
	dp := make([]int, len(nums))
	dp[0] = 0

	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1], nums[i-1]) - 1
		if dp[i] < 0 {
			return false
		}
	}
	return dp[len(nums)-1] >= 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// each spot in DP is the # of spaces left you can move going off
// of your previous jump
// or, put another way, your previous jump's remaining distance

// nums 2 3 1 1 4
//   dp 0 1 2 1 0

// nums 3 2 1 0 4
//   dp 0 2 1 0 -1

// nums 9 2 1 0 0 4
//   dp 0 8 7 6 5 4

// nums 9 2 1 0 0 0 0 0 0 2 4
//   dp 0 8 7 6 5 4 3 2 1 0 1
