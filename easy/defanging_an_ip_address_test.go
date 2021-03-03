package easy

import "testing"

func TestDefangIPAddress(t *testing.T) {
	tests := []struct {
		name string
		input string
		output string
	}{
		{"First", "1.1.1.1", "1[.]1[.]1[.]1"},
		{"Second", "255.100.50.0", "255[.]100[.]50[.]0"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := DefangIPAddress(test.input)
			if result != test.output {
				t.Logf("For %v, expected output should be %v, not %v", test.input, test.output, result)
				t.Fail()
			}
		})
	}
}

