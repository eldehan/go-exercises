package main

func canJump(nums []int) bool {
	lastGoodJump := len(nums) - 1
	for i := lastGoodJump; i >= 0; i-- {
		if i+nums[i] >= lastGoodJump {
			lastGoodJump = i
		}
	}
	return lastGoodJump == 0
}
