package string_algorithms

import (
	"fmt"
)

/*LengthOfTheLongestSubstringWithoutRepeatingCharacters is a function that finds
the length of the longest substring that not have repeating characters in given string
Given a string s, find the length of the longest substring without repeating characters.
*/
func LengthOfTheLongestSubstringWithoutRepeatingCharacters(s string) int {
	fmt.Printf("\nDEBUG FOR String: %s \n", s)
	longestSubstr := struct {
		Start  int
		Finish int
	}{
		Start:  0,
		Finish: 0,
	}
	index := 0
	// charIndexMap keeps the character: index pairs to detect recurrence.
	charIndexMap := make(map[uint8]int, 0)
	for index < len(s) {
		/* if isExist, it means that character already in the substring. So calculate the substring length and
		   continue with next possible substring.
		*/
		if start, isExist := charIndexMap[s[index]]; isExist {
			if (index - start) > (longestSubstr.Finish - longestSubstr.Start) {
				longestSubstr.Start = longestSubstr.Finish
				longestSubstr.Finish = index
			}
			charIndexMap = make(map[uint8]int, 0)
			index = start + 1
		} else {
			charIndexMap[s[index]] = index
			index += 1
		}
	}
	return longestSubstr.Finish - longestSubstr.Start
}

/*StrStr function returns -1 if given haystack string contains needle string,
else returns the first occurrence index of needle string in haystack string.
*/
func StrStr(haystack string, needle string) int {
	haystackLen := len(haystack)
	needleLen := len(needle)

	if haystackLen < needleLen {
		return -1
	}
	index := 0
	for index+needleLen <= haystackLen {
		subStr := haystack[index : index+needleLen]
		if subStr == needle {
			return index
		}
		index += 1
	}
	return -1
}

/*BalancedStringSplit
Balanced strings are those that have an equal quantity of 'L' and 'R' characters.

Given a balanced string s, split it into some number of substrings such that:

Each substring is balanced.
Return the maximum number of balanced strings you can obtain.

Constraints:

	2 <= s.length <= 1000
	s[i] is either 'L' or 'R'.
	s is a balanced string.
*/
func BalancedStringSplit(s string) int {
	r := 0
	l := 0
	count := 0
	for _, char := range s {
		if char == 'R' {
			r += 1
		} else if char == 'L' {
			l += 1
		}
		if r != 0 && l != 0 && r == l {
			count += 1
			r = 0
			l = 0
		}
	}
	return count
}
