// exercise 11-5

package main

import "fmt"

func uniquePaths(rows int, columns int) int {
	if rows == 1 || columns == 1 {
		return 1
	}
	return uniquePaths(rows-1, columns) + uniquePaths(rows, columns-1)
}

func main() {
	fmt.Println(uniquePaths(3, 7))
}
