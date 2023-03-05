package offer

// Node Definition for a Node.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 两次遍历
/*func copyRandomList(head *Node) *Node {
	cur := head
	hair := &Node{}
	tmp := hair
	for cur != nil {
		newNode := &Node{Val: cur.Val, Next: cur.Next}
		cur = cur.Next
		tmp.Next = newNode
		tmp = newNode
	}
	cur = head
	tmp = hair.Next
	for cur != nil {
		rdm := cur.Random
		find := head
		res := hair.Next
		for find != rdm {
			find = find.Next
			res = res.Next
		}
		tmp.Random = res
		cur = cur.Next
		tmp = tmp.Next
	}
	return hair.Next
}*/

// 哈希
/*func copyRandomList(head *Node) *Node {
	hash := make(map[*Node]*Node)
	for cur := head; cur != nil; cur = cur.Next {
		tmp := &Node{Val: cur.Val}
		hash[cur] = tmp
	}
	for cur := head; cur != nil; cur = cur.Next {
		tmp := hash[cur]
		tmp.Next = hash[cur.Next]
		tmp.Random = hash[cur.Random]
	}
	return hash[head]
}*/

// 原链表插入+拆分
func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
	for cur := head; cur != nil; cur = cur.Next.Next {
		cur.Next = &Node{Val: cur.Val, Next: cur.Next}
	}
	for cur := head; cur != nil; cur = cur.Next.Next {
		if cur.Random != nil {
			cur.Next.Random = cur.Random.Next
		}
	}
	cur := head
	ans := head.Next
	for cur.Next != nil {
		next := cur.Next
		cur.Next = cur.Next.Next
		cur = next
	}
	return ans
}
