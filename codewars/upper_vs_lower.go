package main

import (
	"fmt"
	"strings"
)

func UpperVsLower(str string) string {
	upperCount, lowerCount := 0, 0
	for _, r := range str {
		if strings.ToUpper(string(r)) == string(r) {
			upperCount += 1
		} else {
			lowerCount += 1
		}
	}
	if upperCount > lowerCount {
		return strings.ToUpper(str)
	} else {
		return strings.ToLower(str)
	}
}

func main() {
	fmt.Println(UpperVsLower("coDe"))
	fmt.Println(UpperVsLower("CODe"))
	fmt.Println(UpperVsLower("a"))
	fmt.Println(UpperVsLower("Z"))
}
