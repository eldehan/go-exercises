package main

import "fmt"

func countSheep(num int) (result string) {
	for count := 1; count <= num; count++ {
		result = result + fmt.Sprintf("%d sheep...", count)
	}
	return result
}

func main() {
	fmt.Println(countSheep(5))
	fmt.Println(countSheep(0))
	fmt.Println(countSheep(2))
	fmt.Println(countSheep(1))
}
