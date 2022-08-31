// problem - given a string with different brackets, determine if the string is 'valid'
// valid means - open brackets must be closed by the same type of bracket
// open brackets must be closed in the correct order
// every closed bracket must have a corresponding opening bracket of the same time
// closing brackets must follow opening brackets (cannot have a closing bracket first)

// input - string
// output - boolean (true/false)

// string will always be at least 1 char in length
// in this problem, will only consist of brackets
// but let's pretend this isn't the case

// a stack is a good data structure for this
// when we encounter an opening bracket, we add it to the stack
// when we encounter a closing bracket...
// first, check to see if the most recent bracket on the stack matches
// i.e. is the corresponding opening bracket
// if YES pop it off the stack (we have completed a set of brackets)
// if NO then return false - we have out-of-order brackets
// if we reach the end of the input string and the stack still has brackets on it...
// we short some corresponding closing brackets
// return false
// if we reach the end and the stack is empty
// return true!

// how to identify a matching closing bracket
// hash map that functions as a dictionary

// how to identify if a particular character is a bracket, and what type
// regex (one for open brackets, one for closing)

package main

import (
	"regexp"
)

func isValid(s string) bool {
	bracketDict := map[string]string{")": "(", "}": "{", "]": "["}
	stack := []string{}
	opening, _ := regexp.Compile("[({[]")
	closing, _ := regexp.Compile("[)}\\]]")
	for _, r := range s {
		if opening.MatchString(string(r)) {
			stack = append(stack, string(r))
		} else if closing.MatchString(string(r)) {
			if len(stack) == 0 || stack[len(stack)-1] != bracketDict[string(r)] {
				return false
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}

	if len(stack) > 0 {
		return false
	}

	return true
}
