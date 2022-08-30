package main

func minPathSum(grid [][]int) int {
	cols, rows := len(grid), len(grid[0])

	dp := make([][]int, cols)
	for i := 0; i < cols; i++ {
		dp[i] = make([]int, rows)
	}

	for col := 0; col < cols; col++ {
		for row := 0; row < rows; row++ {
			if col == 0 && row == 0 {
				dp[col][row] = grid[col][row]
			} else {
				if col-1 < 0 {
					dp[col][row] = grid[col][row] + dp[col][row-1]
				} else if row-1 < 0 {
					dp[col][row] = grid[col][row] + dp[col-1][row]
				} else {
					dp[col][row] = grid[col][row] + min(dp[col-1][row], dp[col][row-1])
				}
			}
		}
	}

	return dp[cols-1][rows-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
