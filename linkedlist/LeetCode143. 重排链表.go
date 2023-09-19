package linkedlist

func ReorderList(head *ListNode) {
	mid := getMid(head)
	start := mid.Next
	mid.Next = nil
	start = reverse(start)
	merge(head, start)
}

func getMid(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}

func reverse(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func merge(node1 *ListNode, node2 *ListNode) {
	n1 := node1
	n2 := node2
	for n1 != nil && n2 != nil {
		tmp1 := n1.Next
		tmp2 := n2.Next
		n1.Next = n2
		n2.Next = tmp1
		n1 = tmp1
		n2 = tmp2
	}
}
