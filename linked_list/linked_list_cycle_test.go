package linked_list

import "testing"

func TestHasCycle(t *testing.T) {
	first := &ListNode{1, &ListNode{2, nil}}

	node := &ListNode{10, nil}
	nodeB := &ListNode{11, nil}
	nodeC := &ListNode{11, nil}

	node.Next = nodeB
	nodeB.Next = nodeC
	nodeC.Next = node
	second := &ListNode{1, &ListNode{2, node}}
	tests := []struct {
		name   string
		input  *ListNode
		output bool
	}{
		{"First Case", first, false},
		{"Second Case", second, true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HasCycle(test.input)
			if test.output != result {
				t.Logf("For %v, the answer should be %v not %v", test.name, test.output, result)
				t.Fail()
			}
		})
	}
}

