package string_algorithms_test

import (
	"fmt"
	"github.com/eneskzlcn/algorithm/string_algorithms"
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
	for _, test := range testCases {
		length := string_algorithms.LengthOfTheLongestSubstringWithoutRepeatingCharacters(test.Given)
		assert.Equal(t, test.Expected, length, test.Given)
	}
}

func TestGivenHaystackThatContainsNeedleThenItShouldReturnTheFirstOccurrenceIndexWhenStrStrCalled(t *testing.T) {
	type testCase struct {
		Haystack string
		Needle   string
		Expected int
	}
	testCases := []testCase{
		{
			Haystack: "hello",
			Needle:   "ll",
			Expected: 2,
		},
		{
			Haystack: "aaaaa",
			Needle:   "bba",
			Expected: -1,
		},
	}
	for _, test := range testCases {
		result := string_algorithms.StrStr(test.Haystack, test.Needle)
		assert.Equal(t, test.Expected, result, fmt.Sprintf("Haystack: %s, Needle: %s", test.Haystack, test.Needle))
	}
}

func TestBalancedStringSplit(t *testing.T) {
	type testCase struct {
		Given    string
		Expected int
	}
	testCases := []testCase{
		{
			Given:    "RLRRLLRLRL",
			Expected: 4,
		},
		{
			Given:    "RLRRRLLRLL",
			Expected: 2,
		},
		{
			Given:    "LLLLRRRR",
			Expected: 1,
		},
	}
	for _, test := range testCases {
		result := string_algorithms.BalancedStringSplit(test.Given)
		assert.Equal(t, test.Expected, result)
	}
}
