// need to first sort intervals into ascending order

// initialize results array
// append first interval to results array (index 0)
// the current comparing interval is the last interval in results array
// starting with the following (second) interval (index 1)
// compare whether the end of comparing interval (index 1) is greater than the start
// of the next interval
// if so, they should be merged
// replace the end of the comparing interval w/ the end of the next interval
// next interval increments
// if not, append the next interval to the results array
// now the next interval is the comparing interval

package main

import (
	"fmt"
	"sort"
)

func sortIntervals(intervals [][]int) {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
}

func merge(intervals [][]int) [][]int {
	// guard clause; no point wasting time if only 1 item, or 0
	if len(intervals) <= 1 {
		return intervals
	}

	sortIntervals(intervals)

	mergedIntervals := [][]int{}
	mergedIntervals = append(mergedIntervals, intervals[0])

	for _, interval := range intervals[1:] {
		if currentTop := mergedIntervals[len(mergedIntervals)-1]; currentTop[1] >= interval[0] {
			currentTop[1] = interval[1]
		} else {
			mergedIntervals = append(mergedIntervals, interval)
		}
	}
	return mergedIntervals
}

func main() {
	intervals := [][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18},
	}

	intervalsTwo := [][]int{
		{1, 4},
		{4, 5},
	}

	fmt.Println(merge(intervals))
	fmt.Println(merge(intervalsTwo))
}
