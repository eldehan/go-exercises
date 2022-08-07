package kata

import "strings"

func ReverseWords2(str string) string {
	words := strings.Split(str, " ")
	for i, word := range words {
		runes := []rune(word)
		for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		words[i] = string(runes)
	}
	return strings.Join(words, " ")
}

// or better:

func Reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return result
}

func ReverseWords(str string) (result string) {
	words := strings.Split(str, " ")
	for i, v := range words {
		words[i] = Reverse(string(v))
	}

	return strings.Join(words, " ") // reverse those words
}
