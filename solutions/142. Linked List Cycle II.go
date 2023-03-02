package solutions

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func detectCycle(head *ListNode) *ListNode {
	nodeSet := make(map[*ListNode]bool)
	for head != nil {
		if nodeSet[head] {
			return head
		} else {
			nodeSet[head] = true
		}
		head = head.Next
	}
	return nil
}
