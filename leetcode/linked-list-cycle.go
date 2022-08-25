package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// use fast and slow pointer

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slowPointer := head
	fastPointer := head.Next

	for slowPointer != nil && fastPointer != nil {
		if slowPointer == fastPointer {
			return true
		}

		slowPointer = slowPointer.Next
		if fastPointer.Next == nil {
			return false
		} else {
			fastPointer = fastPointer.Next.Next
		}
	}

	return false
}
