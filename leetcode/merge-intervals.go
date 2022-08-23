// need to first sort intervals into ascending order

// initialize results array
// append first interval to results array (index 0)
// the current comparing interval is the last interval in results array
// starting with the following (second) interval (index 1)
// compare whether the end of comparing interval (index 1) is greater than the start
// of the next interval
// if so, they should be merged
// then, check if comparing interval[1] is greater than next interval[1]
// if so, it completely overlaps the next interval
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
		// CHECK FOR TOTAL OVERLAP; IF CURRENTTOP[1] IS > INTERVAL[1]
		if currentTop := mergedIntervals[len(mergedIntervals)-1]; currentTop[1] >= interval[0] {
			// total overlap
			if currentTop[1] > interval[1] {
				continue
			} else {
				currentTop[1] = interval[1]
			}
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

	intervalsThree := [][]int{
		{3, 6},
		{4, 5},
		{7, 10},
		{8, 11},
	}

	fmt.Println(merge(intervals))      // [[1 6] [8 10] [15 18]]
	fmt.Println(merge(intervalsTwo))   // [1, 5]
	fmt.Println(merge(intervalsThree)) // [[3 6] [7 11]]
}
