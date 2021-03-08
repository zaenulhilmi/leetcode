package linked_list

import "testing"

func TestIsPalindrome(t *testing.T) {
	first := ListNode{1, &ListNode{2, &ListNode{2, &ListNode{1, nil}}}}
	second := ListNode{1, &ListNode{2, nil}}
	tests := []struct {
		name   string
		input  *ListNode
		output bool
	}{
		{"first", &first, true},
		{"first", &second, false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IsPalindrome(test.input)

			if result != test.output {
				t.Logf("for %v, the result should be %v not %v", test.input, test.output, result)
				t.Fail()
			}
		})
	}
}

