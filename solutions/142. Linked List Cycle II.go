package solutions

/*func detectCycle(head *ListNode) *ListNode {
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
}*/

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			fast = head
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return fast
		}
	}
	return nil
}
