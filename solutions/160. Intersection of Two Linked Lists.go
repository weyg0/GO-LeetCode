package solutions

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a, b := headA, headB
	for a != b {
		if a == nil {
			a = headB
		}
		if b == nil {
			b = headA
		}
		if a == b {
			break
		}
		a = a.Next
		b = b.Next
	}
	return a
}
