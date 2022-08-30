// create dp matrix that mirrors shape of triangle matrix
// fill bottom row with values (since those are our starting values; nothing to build off yet)
// for each successive level, fill that spot of the dp matrix in with:
// the value in that spot of the triangle plus the smaller of the 2 values below it in the triangle
// continue until you have filled in the dp matrix up to the very top spot
// return the top spot of the dp triangle - this will be the calculated min

package main

func minimumTotal(triangle [][]int) int {
	height := len(triangle)

	dp := make([][]int, height)
	for level := 0; level < height; level++ {
		dp[level] = make([]int, len(triangle[level]))
	}

	for level := height - 1; level >= 0; level-- {
		for index := 0; index <= level; index += 1 {
			if level == height-1 {
				dp[level][index] = triangle[level][index]
			} else {
				dp[level][index] = triangle[level][index] + min(dp[level+1][index], dp[level+1][index+1])
			}
		}
	}

	return dp[0][0]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
