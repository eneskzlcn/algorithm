package string_test

import (
	stringAlgorithms "algorithm/string"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGivenStringThenItShouldReturnTheLengthOfTheLongestSubstringWhenLengthOfTheLongestSubstringWithoutRepeatingCharactersCalled(t *testing.T) {
	type testCase struct {
		Given    string
		Expected int
	}
	testCases := []testCase{
		{
			Given:    "abcabcbb",
			Expected: 3,
		},
		{
			Given:    "bbbbb",
			Expected: 1,
		},
		{
			Given:    "pwwkew",
			Expected: 3,
		},
		{
			Given:    " ",
			Expected: 1,
		},
		{
			Given:    "au",
			Expected: 2,
		},
		{
			Given:    "aab",
			Expected: 2,
		},
	}
	for _, testCase := range testCases {
		length := stringAlgorithms.LengthOfTheLongestSubstringWithoutRepeatingCharacters(testCase.Given)
		assert.Equal(t, testCase.Expected, length, fmt.Sprintf("%s", testCase.Given))
	}
}
