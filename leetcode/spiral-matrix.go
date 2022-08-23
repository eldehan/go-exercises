// initialize results array
// Create and initialize variables of:
// rowstart and colstart
// represent the initial position of row and columns (0 at start)
// rowend and colend
// represent the size of row and column respectively.

// Run outer loop until all the elements are printed
// if length of results array is == length of all elements in matrix
// In each iteration of the outer loop, traverse four inner loops
// left to right (from colStart to colEnd)
// top to bottom (from rowStart + 1 to rowEnd - 1)
// right to left (from colEnd to colStart)
// bottom to top (from rowEnd to rowStart + 1)

package main

func spiralOrder(matrix [][]int) []int {
	result := []int{}
	if len(matrix) == 0 {
		return result
	}

	width, height := len(matrix[0]), len(matrix[0])
	rowStart, rowEnd, colStart, colEnd := 0, height-1, 0, width-1
	allElementsLength := width * height

	for len(result) < allElementsLength {
		for j := colStart; j <= colEnd && len(result) < allElementsLength; j += 1 {
			result = append(result, matrix[rowStart][j])
		}

		for i := rowStart + 1; i <= rowEnd-1 && len(result) < allElementsLength; i += 1 {
			result = append(result, matrix[i][colEnd])
		}

		for j := colEnd; j >= colStart && len(result) < allElementsLength; j -= 1 {
			result = append(result, matrix[rowEnd][j])
		}

		for i := rowEnd - 1; i >= rowStart+1 && len(result) < allElementsLength; i -= 1 {
			result = append(result, matrix[i][colStart])
		}

		colStart += 1
		colEnd -= 1
		rowStart += 1
		rowEnd -= 1
	}

	return result
}

func main() {

}
