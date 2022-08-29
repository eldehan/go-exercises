// base case - reached bottom of triangle (row = len(triangle)-1)
// the minimum path from the top (triangle[0][i]) at index i, is the lesser of the minimum path between
// triangle[row+1][i] and triangle[row+1][i+1]
package main

import "fmt"

func minimumTotal(triangle [][]int) int {
	return minPath(triangle, 0, 0, make(map[string]int))
}

func minPath(triangle [][]int, row, i int, cache map[string]int) int {
	if row == len(triangle) {
		return 0
	}

	key := fmt.Sprintf("%d:%d", row, i)
	if val, ok := cache[key]; ok {
		return val
	}

	leftSum := minPath(triangle, row+1, i, cache)
	rightSum := minPath(triangle, row+1, i+1, cache)
	result := triangle[row][i] + min(leftSum, rightSum)

	cache[key] = result
	return cache[key]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(minimumTotal([][]int{
		{2},
		{3, 4},
		{6, 5, 7},
		{4, 1, 8, 3},
	}))
}

//     2    // minpath will be the lesser of minPath(triangle[row+1][i]) or minPath(triangle[row+1][i+1])
//    3 4
//   6 5 7
//  4 1 8 3
// 9 2 3 7 1
