package string_algorithms_test

import (
	"fmt"
	"github.com/eneskzlcn/algorithm/string_algorithms"
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
		result := string_algorithms.IsValid(test.Given)
		assert.Equal(t, test.Expected, result, fmt.Sprintf("Given string: %s", test.Given))
	}
}

func TestGivenStringThatConsistOfLettersAndSpacesThenItShouldReturnTheLengthOfLastWordThatNotContainSpacesWhenLengthOfLastWordCalled(t *testing.T) {
	type testCase struct {
		Given    string
		Expected int
	}
	testCases := []testCase{
		{
			Given:    "Hello World",
			Expected: 5,
		},
		{
			Given:    "   fly me   to   the moon  ",
			Expected: 4,
		},
		{
			Given:    "luffy is still joyboy",
			Expected: 6,
		},
		{
			Given:    "a ",
			Expected: 1,
		},
	}

	for _, test := range testCases {
		result := string_algorithms.LengthOfLastWord(test.Given)
		assert.Equal(t, test.Expected, result, fmt.Sprintf("Given string: %s", test.Given))
	}
}
