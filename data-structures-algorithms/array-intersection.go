package main

import "fmt"

func findIntersection(arr1, arr2 []int) []int {
	lookup := make(map[int]bool)

	for _, v := range arr1 {
		lookup[v] = true
	}

	resultArr := []int{}

	for _, v := range arr2 {
		if lookup[v] {
			resultArr = append(resultArr, v)
		}
	}

	return resultArr
}

func main() {
	arr1 := []int{1, 2, 3, 4, 5}
	arr2 := []int{0, 2, 4, 6, 8}

	fmt.Println(findIntersection(arr1, arr2))
}
