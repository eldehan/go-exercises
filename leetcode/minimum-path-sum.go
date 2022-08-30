// The path to reach (col, row) must be through one of the 2 cells:(col-1, row) or (col, row-1). So minimum cost to reach (col, row) can be written as “minimum of those 2 cells plus grid[col][row]”.
// minCost(col, row) = min(minCost(col-1, row), minCost(col, row-1)) + grid[col][row]
// base case => reached origin (grid[0][0]) or went out of bounds (col or row is less than 0)

package main

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)

func minPathSum(grid [][]int) int {
	cols, rows := len(grid), len(grid[0])

	// declare and initialize memo matrix filled with -1s
	memo := make([][]int, cols)
	for i := range memo {
		memo[i] = make([]int, rows)
	}
	for _, subarr := range memo {
		for i := range subarr {
			subarr[i] = -1
		}
	}

	return minCost(grid, cols-1, rows-1, memo)
}

func minCost(grid [][]int, col, row int, memo [][]int) int {
	// if we are out of bounds of the matrix, return maxint to ensure the path is not picked
	if (col < 0) || (row < 0) {
		return MaxInt
		// otherwise, if we are at the origin, return the value at that spot
	} else if col == 0 && row == 0 {
		return grid[col][row]
	}

	// if the current spot does not have a saved value in the memo, compute it
	if memo[col][row] == -1 {
		memo[col][row] = grid[col][row] + min(minCost(grid, col-1, row, memo), minCost(grid, col, row-1, memo))
	}

	// return saved memo value
	return memo[col][row]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
