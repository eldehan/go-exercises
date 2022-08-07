// exercise 9-4

package main

import "fmt"

type runeStack struct {
	items []rune
}

func (s *runeStack) Push(i rune) {
	s.items = append(s.items, i)
}

func (s *runeStack) Pop() rune {
	last := len(s.items) - 1

	toRemove := s.items[last]
	s.items = s.items[:last]
	return toRemove
}

func reverseString(str string) string {
	reversingStack := runeStack{}
	result := []rune{}

	for _, letter := range str {
		reversingStack.Push(letter)
	}

	for len(reversingStack.items) > 0 {
		result = append(result, reversingStack.Pop())
	}

	return string(result)
}

func main() {
	fmt.Println(reverseString("abcde")) // edcba
}
