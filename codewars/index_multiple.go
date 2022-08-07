package main

import "fmt"

func multipleOfIndex(ints []int) []int {
	result := []int{}
	for i, v := range ints {
		if i == 0 && v != 0 {
			continue
		}

		if v%i == 0 {
			result = append(result, v)
		}
	}

	return result
}

func main() {
	fmt.Println(multipleOfIndex([]int{22, -6, 32, 82, 9, 25})) // =>  [-6, 32, 25]

	fmt.Println(multipleOfIndex([]int{68, -1, 1, -7, 10, 10})) // => [-1, 10]

	fmt.Println(multipleOfIndex([]int{-56, -85, 72, -26, -14, 76, -27, 72, 35, -21, -67, 87, 0, 21, 59, 27, -92, 68})) // => [-85, 72, 0, 68]
}
