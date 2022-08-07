package main

import (
	"fmt"
	"strconv"
	"strings"
)

func HighAndLow(in string) string {
	ints := strings.Split(in, " ")

	highest, _ := strconv.Atoi(ints[0])
	lowest, _ := strconv.Atoi(ints[0])
	for _, v := range ints {
		val, _ := strconv.Atoi(v)
		if val > highest {
			highest = val
		}
		if val < lowest {
			lowest = val
		}
	}

	return fmt.Sprintf("%d %d", highest, lowest)
}

func main() {
	fmt.Println(HighAndLow("8 3 -5 42 -1 0 0 -9 4 7 4 -4"))
	fmt.Println(HighAndLow("1 2 3"))
}
