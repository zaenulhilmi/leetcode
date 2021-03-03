package easy

import (
	"testing"
)

func TestFindTheHighestAltitude(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output int
	}{
		{"first", []int{-5, 1, 5, 0, -7}, 1},
		{"second", []int{-4, -3, -2, -1, 4, 3, 2}, 0},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := LargestAltitude(test.input)
			if result != test.output {
				t.Logf("for input of %v, the output should be %v not %v", test.input, test.output, result)
				t.Fail()
			}
		})
	}
}

