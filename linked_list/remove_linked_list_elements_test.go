package linked_list

import (
	"reflect"
	"testing"
)

func TestRemoveElements(t *testing.T) {
	third := ListNode{3, nil}
	second := ListNode{2, &third}
	first := ListNode{1, &second}

	seven := ListNode{7, nil}
	seven2 := ListNode{7, &seven}
	seven3 := ListNode{7, &seven2}
	seven4 := ListNode{7, &seven3}

	tests := []struct {
		name        string
		input       *ListNode
		remove      int
		outputArray []int
	}{
		{"First", &first, 2, []int{1, 3}},
		{"Second", &seven4, 7, nil},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := RemoveElements(test.input, test.remove)
			var resultArray []int
			for result != nil {
				resultArray = append(resultArray, result.Val)
				result = result.Next
			}

			if !reflect.DeepEqual(resultArray, test.outputArray) {
				t.Logf("It should be %v but got %v", test.outputArray, resultArray)
				t.Fail()
			}
		})
	}

}

