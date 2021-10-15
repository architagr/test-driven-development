package tutorial2

import (
	"fmt"
	"testing"
)

func TestCheckIfPrime(test *testing.T) {

	input := 2879
	expected := true
	output := CheckIfPrime(input)

	if output != expected {
		test.Error("number {", input, "} CheckIfPrime expected:{", expected, "} but got {", output, "}")
	}
}

func TestCheckIfPrimeTable(test *testing.T) {

	testCases := []struct {
		input    int
		expected bool
	}{
		{2, true},
		{3, true},
		{4, false},
		{7, true},
		{2879, true},
	}

	for _, testCase := range testCases {
		test.Run(fmt.Sprintf("Testing %d", testCase.input), func(t *testing.T) {

			if got := CheckIfPrime(testCase.input); got != testCase.expected {
				t.Errorf("number {%d} CheckIfPrime expected:{%t} but got {%t}", testCase.input, testCase.expected, got)
			}
		})
	}
}
