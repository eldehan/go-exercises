package main

import (
	"fmt"
	"strings"
)

func IsPalindrome(str string) bool {
	var reversedStr string = Reverse(str)
	return strings.ToLower(str) == strings.ToLower(reversedStr)
}

func Reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return result
}

func main() {
	fmt.Println(IsPalindrome("a"))
	fmt.Println(IsPalindrome("aba"))
	fmt.Println(IsPalindrome("Abba"))
	fmt.Println(IsPalindrome("hello"))
}
