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
	fastPointer := head

	for fastPointer != nil && fastPointer.Next != nil {
		slowPointer = slowPointer.Next
		fastPointer = fastPointer.Next.Next

		if slowPointer == fastPointer {
			return true
		}
	}

	return false
}
