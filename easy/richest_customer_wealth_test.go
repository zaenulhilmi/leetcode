package easy

import "testing"

func TestRichestCustomerWealth(t *testing.T) {
	tests := []struct {
		name   string
		input  [][]int
		output int
	}{
		{"First", [][]int{[]int{1, 2, 3}, []int{3, 2, 1}}, 6},
		{"Second", [][]int{[]int{1, 5}, []int{7, 3}, []int{3, 5}}, 10},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := RichestCustomerWealth(test.input)
			if result != test.output {
				t.Logf("For %v, the answer should be %v not %v", test.input, test.output, result)
				t.Fail()
			}
		})
	}
}

