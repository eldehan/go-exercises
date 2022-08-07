package main

import "fmt"

func century(year int) int {
	var century int = 0

	for centuryYears := 0; centuryYears < year; centuryYears += 100 {
		century += 1
	}
	fmt.Println(century)
	return century
}

func main() {
	century(int(1990))
	century(int(1705))
	century(int(1900))
	century(int(1601))
	century(int(2000))
	century(int(89))
}
