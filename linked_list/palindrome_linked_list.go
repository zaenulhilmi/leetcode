package linked_list

func IsPalindrome(head *ListNode) bool {
	current := head
	var box []int
	for current != nil {
		box = append(box, current.Val)
		current = current.Next
	}

	length := len(box)
	for i := 0; i < length; i++ {
		if box[i] != box[length-i-1] {
			return false
		}
	}
	return true
}

