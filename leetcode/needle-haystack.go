// given 2 strings, 'needle' and 'haystack', return INDEX of FIRST occurrence of needle in the haystack, or -1 if needle is not part of haystack
// return 0 if needle is an empty string

// input - 2 strings (needle and haystack)
// output - int (represents the index)

// brute force solution?
// if needle is an empty string, return 0
// iterate character by character through haystack
// save this index as the start index
// compare the char at that index to the first char of the needle string
// if they match, continue comparing the next chars until either
// there's a mismatch between needle and haystack
// the end of the needle string is reached
// if mismatch, break loop and move on to next haystack char as start char
// now that index becomes the new 'start' index
// if end of needle reached, congrats you found the needle
// return the start index
// if end of haystack reached and a match has not been found, return -1

// orrrrr

// above time complexity isn't great (O(M*N) I think)
// instead:
// save length of needle to variable
// iterate through haystack
// at each index, compare substring of haystack from the index through to the index + needle length, against the needle
// first though check that that leaves enough length in the string
// if they match, return starting index
// if not, move on
// time complexity is O(N)? (where N is length of haystack)

package main

import "fmt"

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	needleLen := len(needle)

	for i := 0; i < len(haystack)-needleLen+1; i += 1 {
		if haystack[i:i+needleLen] == needle {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(strStr("hello", "ll"))
}

// Runtime: 1 ms, faster than 51.48% of Go online submissions for Implement strStr().
// Memory Usage: 1.9 MB, less than 93.40% of Go online submissions for Implement strStr().
