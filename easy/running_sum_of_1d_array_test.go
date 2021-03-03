package easy

import (
	"reflect"
	"testing"
)

func TestRunningSum(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output []int
	}{
		{"first", []int{1, 2, 3, 4}, []int{1, 3, 6, 10}},
		{"second", []int{1, 1, 1, 1, 1}, []int{1, 2, 3, 4, 5}},
		{"third", []int{3, 1, 2, 10, 1}, []int{3, 4, 6, 16, 17}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := RunningSum(test.input)
			if !reflect.DeepEqual(result, test.output) {
				t.Logf("For %v, the the expected value is %v not %v", test.input, test.output, result)
				t.Fail()
			}
		})
	}
}

