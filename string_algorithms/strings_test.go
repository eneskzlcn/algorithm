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

func TestGivenTwoBinaryNumberAsStringThenItShouldReturnSumOfTwoNumbersAsStringAndBinaryFormatWhenAddBinaryNumbersCalled(t *testing.T) {
	type testCase struct {
		GivenA   string
		GivenB   string
		Expected string
	}
	testCases := []testCase{
		{
			GivenA:   "11",
			GivenB:   "1",
			Expected: "100",
		},
		{
			GivenA:   "1010",
			GivenB:   "1011",
			Expected: "10101",
		},
		{
			GivenA:   "10100000100100110110010000010101111011011001101110111111111101000000101111001110001111100001101",
			GivenB:   "110101001011101110001111100110001010100001101011101010000011011011001011101111001100000011011110011",
			Expected: "110111101100010011000101110110100000011101000101011001000011011000001100011110011010010011000000000",
		},
	}

	for _, test := range testCases {
		result := string_algorithms.AddBinary2(test.GivenA, test.GivenB)
		assert.Equal(t, test.Expected, result, fmt.Sprintf("Given A: %s, Given B: %s", test.GivenA, test.GivenB))
	}
}

func TestGivenAlphanumericAndNonalphanumericMixedStringThenItShouldCleanNonalphanumericsAndChecksIfIsPalindromeWhenIsPalindromeCalled(t *testing.T) {
	type testCase struct {
		Given    string
		Expected bool
	}
	testCases := []testCase{
		{
			Given:    "A man, a plan, a canal: Panama",
			Expected: true,
		},
		{
			Given:    "race a car",
			Expected: false,
		},
		{
			Given:    " ",
			Expected: true,
		},
	}
	for _, test := range testCases {
		result := string_algorithms.IsPalindrome(test.Given)
		assert.Equal(t, test.Expected, result, fmt.Sprintf("Given string: %s", test.Given))
	}
}
