package linked_list

import (
	"reflect"
	"testing"
)

func TestToArray(t *testing.T) {
	head := &ListNode{1, &ListNode{2, &ListNode{3, nil}}}

	result := head.ToArray()

	if !reflect.DeepEqual(result, []int{1,2,3}) {
		t.Logf("The array should be 1,2,3 not %v", result)
	}
}
