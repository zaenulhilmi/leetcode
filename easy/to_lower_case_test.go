package easy

import (
	"testing"
)

func TestToLowerCase(t *testing.T) {
	tests := []struct {
		name string
		input string
		output string
	}{
		{"First", "Hello", "hello"},
		{"Second", "DOTA", "dota"},
		{"Second", "2nd Life", "2nd life"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := ToLowerCase(test.input)
			if result != test.output {
				t.Logf("For %v, the output should be %v, not %v", test.input, test.output, result)
				t.Fail()
			}
		})
	}
}

