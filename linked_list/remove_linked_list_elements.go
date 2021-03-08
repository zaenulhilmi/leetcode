package linked_list

func RemoveElements(head *ListNode, val int) *ListNode {
	current := head
	var prev *ListNode

	for current != nil {
		if val == current.Val {
			if prev == nil {
				head = current.Next
			} else {
				prev.Next = current.Next
			}
		} else {
			prev = current
		}
		current = current.Next
	}

	return head
}
