package string_test

import (
	stringAlgorithms "algorithm/string"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGivenStringThatContainsJustOpeningAndClosingBracketsThatClosesInCorrectOrderThenItShouldReturnTrueWhenIsValidCalled(t *testing.T) {
	type testCase struct {
		Given    string
		Expected bool
	}
	testCases := []testCase{
		{
			Given:    "()",
			Expected: true,
		},
		{
			Given:    "()[]{}",
			Expected: true,
		},
		{
			Given:    "(}",
			Expected: false,
		},
		{
			Given:    "[",
			Expected: false,
		},
	}
	for _, test := range testCases {
		result := stringAlgorithms.IsValid(test.Given)
		assert.Equal(t, test.Expected, result, fmt.Sprintf("Given string: %s", test.Given))
	}
}
