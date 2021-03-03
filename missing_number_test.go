package main

import "testing"

func TestMissingNumber(t *testing.T) {
	tests := []struct {
		name string
		input []int
		output int
	}{
		{"first", []int{3,0,1}, 2},
		{"second", []int{0,1}, 2},
		{"third", []int{9,6,4,2,3,5,7,0,1}, 8},
		{"fourth", []int{0}, 1},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			result := MissingNumber(test.input)

			if result != test.output {
				t.Logf("From %v the answer should be %v not %v", test.input, test.output, result)
				t.Fail()
			}
		})
	}
}

