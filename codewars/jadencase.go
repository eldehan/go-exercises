package main

import (
	"fmt"
	"strings"
	"unicode"
)

func CapitalizeWord(str string) string {
	runes := []rune(str)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func ToJadenCase(str string) string {
	splitStr := strings.Split(str, " ")
	for i, word := range splitStr {
		splitStr[i] = CapitalizeWord(word)
	}

	return strings.Join(splitStr, " ")
}

func main() {
	fmt.Println(ToJadenCase("most trees are blue"))
	fmt.Println(ToJadenCase("All the rules in this world were made by someone no smarter than you. So make your own."))
	fmt.Println(ToJadenCase("When I die. then you will realize"))
	fmt.Println(ToJadenCase("Jonah Hill is a genius"))
	fmt.Println(ToJadenCase("Dying is mainstream"))
}
