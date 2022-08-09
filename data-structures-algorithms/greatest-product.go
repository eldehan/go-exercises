// exercise 13-1

// given an array of positive numbers, write a function
// that returns the greatest product of any 3 numbers.
// efficiency should be O(N log N)

package main

import (
	"fmt"
	"sort"
)

func greatestProduct(numbers []int) int {
	sort.Ints(numbers)
	threeNums := numbers[len(numbers)-3:]

	product := 1

	for _, v := range threeNums {
		product *= v
	}

	return product
}

func main() {
	fmt.Println(greatestProduct([]int{1, 7, 3, 2, 4})) // 7 * 3 * 4 == 84
}
