package kata

import "fmt"

// total = len of string
// # of errors = # of letters NOT from a-m
// initialize # of errors var
// loop thru string
// check if rune falls between a-m
// if not, add 1 to error count
// return string of errors/len

func PrinterError(s string) string {
	var strLen int = len(s)
	var errCount int = 0

	for _, r := range s {
		if !(r >= 'a') || !(r <= 'm') {
			errCount += 1
		}
	}

	return fmt.Sprintf("%d/%d", errCount, strLen)
}
