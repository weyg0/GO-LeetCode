package offers

import "../solutions"

type ListNode = solutions.ListNode

func reversePrint(head *ListNode) []int {
	var ans []int
	for head != nil {
		ans = append(ans, head.Val)
		head = head.Next
	}
	n := len(ans)
	for i := 0; i < n/2; i++ {
		tmp := ans[i]
		ans[i] = ans[n-i-1]
		ans[n-i-1] = tmp
	}
	return ans
}
