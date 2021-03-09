package linked_list

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func (head ListNode) ToArray() interface{} {
	fmt.Println(head)
	current := &head
	var result []int
	for current != nil {
		result = append(result, current.Val)
		current = current.Next
	}
	return result
}
