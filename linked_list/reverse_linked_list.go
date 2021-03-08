package linked_list

func ReverseList(head *ListNode) *ListNode {
	current := head
	var prev, next *ListNode

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	return prev
}
