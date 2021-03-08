package linked_list

import (
	"reflect"
	"testing"
)

func TestReverseList(t *testing.T) {
	tests := []struct {
		name        string
		input       []int
		outputArray []int
	}{
		{"First", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{"Second", []int{1, 2}, []int{2, 1}},
		{"Third", []int{}, []int{}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			var head, prev *ListNode
			for _, v := range test.input {
				current := ListNode{v, nil}
				if head == nil {
					head = &current
				}
				if prev != nil {
					prev.Next = &current
				}
				prev = &current
			}
			result := ReverseList(head)

			resultArray := []int{}

			for result != nil {
				resultArray = append(resultArray, result.Val)
				result = result.Next
			}

			if !reflect.DeepEqual(resultArray, test.outputArray) {
				t.Logf("It should be %v not %v", test.outputArray, resultArray)
				t.Fail()
			}
		})
	}
}

