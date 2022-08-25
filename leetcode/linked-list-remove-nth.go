package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// iterate thru list and count how many nodes there are
// keep going as long as the next node is not nil (if it's nil then you're on the last element)

// then, reset currentNode to the dummy head
// now, iterate thru list as long as the counter is < counter - n
// this will be the node before the nth node from the end
// delete that node
// return head

// func removeNthFromEnd(head *ListNode, n int) *ListNode {
//   dummyHead := &ListNode{Next: head}
// 	currentNode := dummyHead
//   listLength := 0

//   for currentNode.Next != nil {
//     currentNode = currentNode.Next
//     listLength += 1
//   }

//   currentNode = dummyHead
//   for i := 0; i < listLength - n; i += 1 {
//     currentNode = currentNode.Next
//   }

//   currentNode.Next = currentNode.Next.Next

//   return dummyHead.Next
// }

// alternate solution that only requires 1 pass

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy
	target := head
	nilFinder := head

	for i := 0; i < n; i++ {
		nilFinder = nilFinder.Next
	}

	for nilFinder != nil {
		prev = target
		target = target.Next
		nilFinder = nilFinder.Next
	}

	prev.Next = target.Next

	return dummy.Next
}
