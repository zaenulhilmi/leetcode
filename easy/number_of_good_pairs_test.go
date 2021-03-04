package easy

import "testing"

func TestNumIdenticalPairs(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output int
	}{
		{"First", []int{1, 2, 3, 1, 1, 3}, 4},
		{"Second", []int{1, 1, 1, 1}, 6},
		{"Third", []int{1, 2, 3}, 0},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := NumIdenticalPairs(test.input)
			if result != test.output {
				t.Logf("For %v the answer should be %v not %v", test.input, test.output, result)
				t.Fail()
			}
		})
	}
}

