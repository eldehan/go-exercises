// Question:
// Given a string, determine if it is a palindrome, considering only alphanumeric characters
// and ignoring cases.
// For example,
// "A man, a plan, a canal: Panama" is a palindrome.
// "race a car" is not a palindrome

// an empty string is a palindrome

// 2 pointers, 1 at the beginning of the string and one at the end
// pointers move towards each other, comparing characters until they meet
// ignore non alphanumerics
// if pointers meet and all match, then it's a palindrome
// if at any point the chars do not match, it is not a palindrome

// time complexity of O(N), space complexity of O(1)

package main

import (
	"fmt"
	"strings"
	"unicode"
)

func validPalindrome(str string) bool {
	str = strings.ToLower(str)

	for i, j := 0, len(str)-1; i <= j; {
		for !isAlphanumeric(rune(str[i])) {
			i++
		}

		for !isAlphanumeric(rune(str[j])) {
			j--
		}

		if str[i] != str[j] {
			return false
		}

		i++
		j--
	}

	return true
}

func isAlphanumeric(char rune) bool {
	if unicode.IsLetter(char) || unicode.IsNumber(char) {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(validPalindrome("tot"))
	fmt.Println(validPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(validPalindrome("race a car"))
}
