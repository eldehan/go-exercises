package kata

func Solution(word string) string {
	runeSlice := []rune(word) // making slice of word, so that each element is a rune
	for i, j := 0, len(runeSlice)-1; i < j; i, j = i+1, j-1 {
		runeSlice[i], runeSlice[j] = runeSlice[j], runeSlice[i]
	}
	return string(runeSlice)
}

// or

func Solution2(word string) (result string) {
	for _, value := range word {
		result = string(value) + result
	}
	return
}
