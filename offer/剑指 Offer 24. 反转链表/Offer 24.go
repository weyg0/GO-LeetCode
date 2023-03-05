package offer

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var tail = &ListNode{}
	tail = nil
	for head != nil {
		next := head.Next
		head.Next = tail
		tail = head
		head = next
	}
	return tail
}
