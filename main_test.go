package main

import (
	"reflect"
	"testing"
)

func TestFindErrorNums(t *testing.T) {
	tests := []struct {
		name   string
		input  []int
		output []int
	}{
		{"First", []int{1, 2, 2, 4}, []int{2, 3}},
		{"Second", []int{1,1}, []int{1,2}},
		{"third", []int{3,2,2}, []int{2, 1}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := FindErrorNums(test.input)
			if !reflect.DeepEqual(result, test.output) {
				t.Logf("for %v, expecting %v but got %v", test.input, test.output, result)
				t.Fail()
			}
		})
	}
}

func FindErrorNums(input []int) []int {
	var first, second int
	dict := map[int]int{}
	for _, num := range input {
		if dict[num] == 0 {
			dict[num] = 1
		} else {
			first = num
		}
	}

	for i, _ := range input {
		if dict[i+1] == 0 {
			second = i+1
		}
	}

	return []int{first, second}
}
