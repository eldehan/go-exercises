package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// make dummy node which points to head node
// starting with dummy node, check if next node's value is NOT null
// if so, check if its equal to target
// if = to target, current node's next is set to point to next's next

func removeElements(head *ListNode, targetVal int) *ListNode {
	dummyHead := &ListNode{Next: head}
	currentNode := dummyHead

	for currentNode.Next != nil {
		if currentNode.Next.Val == targetVal {
			currentNode.Next = currentNode.Next.Next
		} else {
			currentNode = currentNode.Next
		}
	}

	return dummyHead.Next
}
