package linked_list

import (
	"reflect"
	"testing"
)

func TestDeleteNode(t *testing.T) {
	second := &ListNode{2, &ListNode{3, &ListNode{4, nil}}}
	head := &ListNode{1, second}

	DeleteNode(second)
	if !reflect.DeepEqual(toArray(head), []int{1, 3, 4}) {
		t.Logf("should be 1,3,4 not %v", toArray(head))
		t.Fail()
	}

}

func toArray(node *ListNode) []int {
	var result []int
	for node != nil {
		result = append(result, node.Val)
		node = node.Next
	}
	return result
}
