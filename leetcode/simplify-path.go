// simplify an absolute path down to a canonical path
// 	eliminate duplicate slashes (i.e. //)
// 	no trailing slash at end of path
// 	if path is '/./' you just stay in the same directory
// 	if path is '/../' you go up a level (to containing directory)
//	otherwise, if path is a name, you are moving down into that directory

// input - string
// output - reduced string

// stack is a good candidate for this (mental model, maybe think of it as a horizontal stack)
// split string into array of strings to iterate through
// 	split based on '/' char
// declare and initialize a new slice to serve as the stack
// iterate through slice of split strings
// 	if the string is '.' or an empty string (indicating we split a double slash '//')
// 		do nothing; don't add it to stack, don't pop off; just continue iterating
// 	if string is '..'
// 		we are moving up a level (assuming there are levels to move up into)
// 		if there is something on the stack, pop it off
// 	otherwise, if the string is anything else, we interpret it as a directory we are moving into
// 		add it to the stack
// when we reach the end, add a "/" to the beginning and join the stack together into a new string
// return new reduced string

package main

import "strings"

func simplifyPath(path string) string {
	files := strings.Split(path, "/")

	reduced := make([]string, 0, len(files))

	for _, path := range files {
		length := len(reduced)
		if path == "." || path == "" {
			continue
		} else if path == ".." {
			if length >= 1 {
				reduced = reduced[0 : length-1]
			}
		} else {
			reduced = append(reduced, path)
		}
	}

	return "/" + strings.Join(reduced, "/")
}
