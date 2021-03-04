package easy

import (
	"reflect"
	"testing"
)

func TestShuffleTheArray(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output []int
	}{
		{"First", []int{2, 5, 1, 3, 4, 7}, []int{2, 3, 5, 4, 1, 7}},
		{"Second", []int{1, 2, 3, 4, 4, 3, 2, 1}, []int{1, 4, 2, 3, 3, 2, 4, 1}},
		{"Third", []int{1, 1, 2, 2}, []int{1, 2, 1, 2}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ShuffleTheArray(test.input)
			if !reflect.DeepEqual(result, test.output) {
				t.Logf("For %v, the answer should be %v not %v", test.input, test.output, result)
				t.Fail()
			}
		})
	}
}

