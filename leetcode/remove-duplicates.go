// input - array of integers sorted in non-decreasing order
// goal - remove duplicates in place
// output - array of integers where first part of array is the non-duplicates, then after that is the duplicates
// do not allocate extra space for another array
// must modify input array in place with O(1) extra memory

// pointers!...

// size (k) = 1
// current item = a
//  a r                - increment r until val != a
// [0,0,1,1,1,2,2,3,3,4]
//  a   r               - first num that is not a; swap vals of a+1 and r; k+= 1
// [0,0,1,1,1,2,2,3,3,4]
//  a   r
// [0,1,0,1,1,2,2,3,3,4]
//    a   r             - next, increment a and r
// [0,1,0,1,1,2,2,3,3,4]
//    a       r         - increment r until val != a
// [0,1,0,1,1,2,2,3,3,4]
//    a       r         - swap vals of a+1 and r; k+= 1
// [0,1,2,1,1,0,2,3,3,4]
//    a       r
// [0,1,2,1,1,0,2,3,3,4]
//      a       r       - increment a and r
// [0,1,2,1,1,0,2,3,3,4]
//      a         r     - increment r until val != a
// [0,1,2,1,1,0,2,3,3,4]
//      a         r     - swap a+1 and r; k+= 1
// [0,1,2,3,1,0,2,2,3,4]
//        a         r
// [0,1,2,3,1,0,2,2,3,4]
//        a           r
// [0,1,2,3,1,0,2,2,3,4]
//        a           r swap; k += 1
// [0,1,2,3,4,0,2,2,3,1]

// so, in the end, k = 5

// algo summarized:
// - k starts at 1 (representing first item of array)
// - a starts at array[0], r starts at array[1]
// - increment r until the value of array[r] != the value of array[a]
//   - at this point, swap vals of a+1 and r
//   - increment k
//   - increment a and r
// - loop ends when r goes out of bounds of array

// optimizing note: a will mark the index of the final spot of unique elments; so, can just return a+1 instead of tracking
// a separate k variable

package main

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	a, r := 0, 1

	for r <= len(nums)-1 {
		if nums[r] == nums[a] {
			r++
		} else {
			nums[r], nums[a+1] = nums[a+1], nums[r]
			a++
			r++
		}
	}

	return a + 1
}
