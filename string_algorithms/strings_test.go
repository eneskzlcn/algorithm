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
func TestGivenColumnNumberThenItShouldReturnCorrespendingNumberInEnglishAlphabetBaseNumberSystemWhenConvertToTitleCalled(t *testing.T) {
	type testCase struct {
		Given    int
		Expected string
	}
	testCases := []testCase{
		{
			Given:    1,
			Expected: "A",
		},
		{
			Given:    26,
			Expected: "Z",
		},
		{
			Given:    25,
			Expected: "Y",
		},
		{
			Given:    27,
			Expected: "AA",
		},
		{
			Given:    28,
			Expected: "AB",
		},
		{
			Given:    701,
			Expected: "ZY",
		},
	}
	for _, test := range testCases {
		result := string_algorithms.ConvertToTitle(test.Given)
		assert.Equal(t, test.Expected, result, fmt.Sprintf("Given number: %d", test.Given))
	}
}

func TestIsIsomorphic(t *testing.T) {
	type testCase struct {
		GivenS   string
		GivenT   string
		Expected bool
	}
	t.Run("Given two non-isomorphic strings then it should return false when IsIsomorphic called.", func(t *testing.T) {
		testCases := []testCase{
			{
				GivenS:   "foo",
				GivenT:   "bar",
				Expected: false,
			},
			{
				GivenS:   "fasaa",
				GivenT:   "blmkk",
				Expected: false,
			},
		}
		for _, test := range testCases {
			result := string_algorithms.IsIsomorphic(test.GivenS, test.GivenT)
			assert.Equal(t, test.Expected, result, fmt.Sprintf("Given S: %s, Given T: %s", test.GivenS, test.GivenT))
		}
	})
	t.Run("Given two isomorphic strings then it should return true when IsIsomorphic called.", func(t *testing.T) {
		testCases := []testCase{
			{
				GivenS:   "egg",
				GivenT:   "add",
				Expected: true,
			},
			{
				GivenS:   "paper",
				GivenT:   "title",
				Expected: true,
			},
		}
		for _, test := range testCases {
			result := string_algorithms.IsIsomorphic(test.GivenS, test.GivenT)
			assert.Equal(t, test.Expected, result, fmt.Sprintf("Given S: %s, Given T: %s", test.GivenS, test.GivenT))
		}
	})
}

func TestIsAnagram(t *testing.T) {
	type testCase struct {
		GivenS   string
		GivenT   string
		Expected bool
	}
	testCases := []testCase{
		{
			GivenS:   "anagram",
			GivenT:   "nagaram",
			Expected: true,
		},
		{
			GivenS:   "word",
			GivenT:   "drow",
			Expected: true,
		},
		{
			GivenS:   "rat",
			GivenT:   "car",
			Expected: false,
		},
	}
	for _, test := range testCases {
		result := string_algorithms.IsAnagram(test.GivenS, test.GivenT)
		assert.Equal(t, test.Expected, result, fmt.Sprintf("Given S: %s, Given T: %s", test.GivenS, test.GivenT))
	}
}
func TestDefangIPaddr(t *testing.T) {
	type testCase struct {
		Given    string
		Expected string
	}
	testCases := []testCase{
		{
			Given:    "1.1.1.1",
			Expected: "1[.]1[.]1[.]1",
		},
		{
			Given:    "255.100.50.0",
			Expected: "255[.]100[.]50[.]0",
		},
	}
	for _, test := range testCases {
		result := string_algorithms.DefangIPaddr(test.Given)
		assert.Equal(t, test.Expected, result)
	}
}
func TestFinalValueAfterOperations(t *testing.T) {
	type testCase struct {
		Given    []string
		Expected int
	}
	testCases := []testCase{
		{
			Given:    []string{"--X", "X++", "X++"},
			Expected: 1,
		},
		{
			Given:    []string{"++X", "++X", "X++"},
			Expected: 3,
		},
		{
			Given:    []string{"X++", "++X", "--X+", "X--"},
			Expected: 0,
		},
	}
	for _, test := range testCases {
		result := string_algorithms.FinalValueAfterOperations(test.Given)
		assert.Equal(t, test.Expected, result)
	}
}
func TestMinPartitions(t *testing.T) {
	type testCase struct {
		Given    string
		Expected int
	}
	testCases := []testCase{
		{
			Given:    "32",
			Expected: 3,
		},
		{
			Given:    "82734",
			Expected: 8,
		},
		{
			Given:    "27346209830709182346",
			Expected: 9,
		},
	}
	for _, test := range testCases {
		result := string_algorithms.MinPartitions(test.Given)
		assert.Equal(t, test.Expected, result)
	}
}

func TestMinOperations(t *testing.T) {
	type testCase struct {
		Given    string
		Expected []int
	}
	testCases := []testCase{
		{
			Given:    "110",
			Expected: []int{1, 1, 3},
		},
		{
			Given:    "001011",
			Expected: []int{11, 8, 5, 4, 3, 4},
		},
	}
	for _, test := range testCases {
		result := string_algorithms.MinOperations(test.Given)
		assert.ElementsMatch(t, test.Expected, result)
	}
}

func TestExecuteInstructions(t *testing.T) {
	type testCase struct {
		GivenN        int
		GivenS        string
		GivenStartPos []int
		Expected      []int
	}
	testCases := []testCase{
		{
			GivenN:        3,
			GivenS:        "RRDDLU",
			GivenStartPos: []int{0, 1},
			Expected:      []int{1, 5, 4, 3, 1, 0},
		},
		{
			GivenN:        2,
			GivenS:        "LURD",
			GivenStartPos: []int{1, 1},
			Expected:      []int{4, 1, 0, 0},
		},
	}
	for _, test := range testCases {
		result := string_algorithms.ExecuteInstructions(test.GivenN, test.GivenStartPos, test.GivenS)
		assert.ElementsMatch(t, test.Expected, result)
	}
}

func TestMostWordsFound(t *testing.T) {
	type testCase struct {
		Given    []string
		Expected int
	}
	testCases := []testCase{
		{
			Given:    []string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"},
			Expected: 6,
		},
		{
			Given:    []string{"please wait", "continue to fight", "continue to win"},
			Expected: 3,
		},
	}
	for _, test := range testCases {
		result := string_algorithms.MostWordsFound(test.Given)
		assert.Equal(t, test.Expected, result)
	}
}
