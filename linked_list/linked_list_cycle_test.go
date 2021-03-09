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

func TestDetectCycle(t *testing.T) {
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
		output *ListNode
	}{
		{"First Case", first, nil},
		{"Second Case", second, node},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := DetectCycle(test.input)
			if test.output != result {
				t.Logf("For %v, the answer should be %v not %v", test.name, test.output, result)
				t.Fail()
			}

			result2 := DetectCycle2(test.input)
			if test.output != result2 {
				t.Logf("For %v, the answer should be %v not %v", test.name, test.output, result2)
				t.Fail()
			}
		})
	}
}

func DetectCycle2(head *ListNode) *ListNode {
	var current *ListNode
	current = head
	intersection := hasCycle(current)

	if intersection == nil {
		return nil
	}

	current = head
	for current != nil {
		if intersection == current {
			return current
		}
		intersection = intersection.Next
		current = current.Next
	}
	return nil
}

func hasCycle(current *ListNode) *ListNode {
	if current == nil {
		return nil
	}
	var slow, fast *ListNode
	slow = current
	fast = current

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return slow
		}
	}

	return nil
}

func DetectCycle(head *ListNode) *ListNode {
	box := make(map[*ListNode]bool)
	current := head
	for current != nil {
		if box[current] {
			return current
		}
		box[current] = true
		current = current.Next
	}
	return nil
}
