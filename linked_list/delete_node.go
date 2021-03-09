package linked_list

func DeleteNode(node *ListNode) {
	if node == nil {
		return
	}
	if node.Next != nil {
		node.Val = node.Next.Val
		node.Next = node.Next.Next
	}
}

