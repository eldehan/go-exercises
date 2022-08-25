// algorithm must be O(log n) complexity...
// binary search
// grab mid
// mid < target value?
// start = mid+1
// mid > target value?
// end = mid-1
// mid == target value?
// now move w/ pointers
// both start on mid
// while prev value is also target value, move start pointer down
// while next value is target value, move end pointer up
// once next and prev are no longer target value, return indexes

// BINARY SEARCH WITH MOVING POINTERS (O(NlogN) worst case)
// func searchRange(nums []int, target int) []int {
//   result := []int{-1, -1}
// 	left, right := 0, len(nums)-1

// 	for left <= right {
// 		mid := left + ((right - left) / 2)
// 		if nums[mid] == target {
//       firstOccurence := mid
//       lastOccurence := mid

//       for firstOccurence-1 >= left && nums[firstOccurence-1] == target {
//         firstOccurence -= 1
//       }

//       for lastOccurence+1 <= right && nums[lastOccurence+1] == target {
//         lastOccurence += 1
//       }

//       result = []int{firstOccurence, lastOccurence}
// 			break
// 		} else if nums[mid] < target {
// 			left = mid + 1
// 		} else {
// 			right = mid - 1
// 		}
// 	}
// 	return result
// }

// BINARY SEARCH ONLY (O(2logN), simplified to O(logN))
//  l         m          r
// [1,5,7,7,7,8,8,8,9,10,11] - i=0; result = [5, -1]

//  l   m   r
// [1,5,7,7,7,8,8,8,9,10,11] - no match, will keep running binary searches until left is not <= right

// then, i will be 1

//  l         m          r
// [1,5,7,7,7,8,8,8,9,10,11] - i=1; result = [5, 5]

//	l   m    r
//
// [1,5,7,7,7,8,8,8,9,10,11] - mid no longer target, but > than; right adjusted
//
//	ml r
//
// [1,5,7,7,7,8,8,8,9,10,11] - result now [5, 6]; mid = target again, i = 1 still, so now left = mid+1
//
//	lmr
//
// [1,5,7,7,7,8,8,8,9,10,11] - left mid right all index 7 now; result = [5, 7] which is our answer
package main

func searchRange(nums []int, target int) []int {
	result := []int{-1, -1}

	for i := 0; i < 2; i++ {
		left, right := 0, len(nums)-1
		for left <= right {
			mid := left + ((right - left) / 2)
			if nums[mid] == target {
				if i == 0 {
					// i == 0 -> locate first occurrence; first run of overall loop
					// as long as left is less than or equal to right and we keep getting matches,
					// the first occurence will be assigned to the current match (mid), and the right (ending)
					// boundary will be moved up to 1 before mid
					// therefore, running another binary search on remaining lower half
					result[0] = mid
					right = mid - 1
				} else {
					// i == 1 -> locate final occurence; second run of overall loop
					// as long as left is less than or equal to right and we keep getting matches,
					// the last occurence will be assigned to the current match (mid), and the left (starting)
					// boundary will be moved up to 1 after mid
					// therefore, running another binary search on remaining upper half
					result[1] = mid
					left = mid + 1
				}
			} else if nums[mid] < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return result
}
