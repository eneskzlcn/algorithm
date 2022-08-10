package string_algorithms

import (
	"fmt"
	stringUtils "github.com/eneskzlcn/utilities/string"
	"strconv"
	"strings"
)

func removeLast(slice []int32) []int32 {
	if len(slice) > 0 {
		slice = slice[:len(slice)-1]
	}
	return slice
}

const (
	OPEN_BRACKET         int32 = 40
	CLOSE_BRACKET        int32 = 41
	OPEN_SQUARE_BRACKET  int32 = 91
	CLOSE_SQUARE_BRACKET int32 = 93
	OPEN_CURLY_BRACKET   int32 = 123
	CLOSE_CURLY_BRACKET  int32 = 125
)

var bracketsMap = map[int32]int32{
	CLOSE_BRACKET:        OPEN_BRACKET,
	CLOSE_SQUARE_BRACKET: OPEN_SQUARE_BRACKET,
	CLOSE_CURLY_BRACKET:  OPEN_CURLY_BRACKET,
}

func IsValid(s string) bool {
	stack := make([]int32, 0)
	for _, char := range s {
		if char == CLOSE_BRACKET || char == CLOSE_SQUARE_BRACKET || char == CLOSE_CURLY_BRACKET {
			if len(stack) <= 0 {
				return false
			}
			lastElement := stack[len(stack)-1]
			if value, ok := bracketsMap[char]; ok {
				if lastElement == value {
					stack = removeLast(stack)
				} else {
					return false
				}
			}
		} else {
			stack = append(stack, char)
		}
	}
	return len(stack) <= 0
}

/*LengthOfLastWord takes a string as param that consist of spaces or letters only.
It returns length of the last words (letters without spaces)
*/
func LengthOfLastWord(s string) int {
	const whiteSpace uint8 = 32
	lastIndex := len(s) - 1
	for lastIndex >= 0 {
		if s[lastIndex] == whiteSpace {
			lastIndex -= 1
		} else {
			break
		}
	}
	if lastIndex <= 0 {
		return 1
	}
	firstIndex := lastIndex - 1
	for firstIndex >= 0 {
		if s[firstIndex] != whiteSpace {
			firstIndex -= 1
		} else {
			return lastIndex - firstIndex
		}
	}
	return lastIndex - firstIndex
}

/*AddBinary takes two strings as binary numbers like "1010", "1011" and
returns their sum as binary string "10101"
*/
//func AddBinary(a string, b string) string {
//	//overflow := 0
//	//
//	//aIndex := len(a) - 1
//	//bIndex := len(b) - 1
//	//step := 0
//	//var binaryMapper = map[uint8]int{
//	//	48: 0, 49: 1,
//	//}
//	//result := ""
//	//for aIndex >= 0 || bIndex >= 0 {
//	//	bValue := 0
//	//	aValue := 0
//	//	if bIndex >= 0 {
//	//		bValue = binaryMapper[b[bIndex]] * int(math.Pow(2, float64(step)))
//	//	}
//	//	if aIndex >= 0 {
//	//		aValue = binaryMapper[a[aIndex]] * int(math.Pow(2, float64(step)))
//	//	}
//	//	bit := aValue + bValue + overflow
//	//
//	//	result := fmt.Sprintf("%d%s", bit, result)
//	//	overflow -= 1
//	//}
//	//return ""
//	return ""
//}

/*AddBinary2 takes two strings as binary numbers like "1010", "1011" and
returns their sum as binary string "10101"
But not works for more than 64 bits.
*/
func AddBinary2(a string, b string) string {
	fmt.Printf("A len %d", len(a))
	aValue, _ := strconv.ParseInt(a, 2, 64)
	bValue, _ := strconv.ParseInt(b, 2, 64)
	result := strconv.FormatInt(aValue+bValue, 2)
	return result
}

func IsPalindrome(s string) bool {
	clearedS := stringUtils.StripNonAlphanumerics(s)
	clearedS = strings.ToLower(clearedS)
	if len(clearedS) == 0 { // accepts "" as palindrome
		return true
	}
	start := 0
	end := len(clearedS) - 1
	for start < end {
		if clearedS[start] != clearedS[end] {
			return false
		}
		end--
		start++
	}
	return true
}

/*ConvertToTitle
Given an integer columnNumber, return its corresponding column title as it appears in an Excel sheet.
For example:
A -> 1
B -> 2
C -> 3
Z -> 26
AA -> 27
AB -> 28
*/
func ConvertToTitle(columnNumber int) string {
	//ASCII CHAR 'A' DECIMAL VALUE = 65
	value := columnNumber
	result := ""
	for value > 0 {
		bit := (value - 1) % 26 // 1
		value = value - bit
		value = (value - 1) / 26
		result = fmt.Sprintf("%c", int32(bit+65)) + result
	}
	return result
}

/*IsIsomorphic Given two strings s and t, determine if they are isomorphic.

Two strings s and t are isomorphic if the characters in s can be replaced to get t.

All occurrences of a character must be replaced with another character while preserving
the order of characters. No two characters may map to the same character, but a character may map to itself.

Constraints:
- 1 <= s.length <= 5 * 104
- t.length == s.length
- s and t consist of any valid ascii character.

From https://leetcode.com/problems/isomorphic-strings/
*/
func IsIsomorphic(s string, t string) bool {
	charTransformationMap := make(map[int32]int32, 0)
	for index, char := range s {
		toValue := int32(t[index])
		if value, isExist := charTransformationMap[char]; isExist {
			if value != toValue {
				return false
			}
		} else {
			for _, val := range charTransformationMap {
				if val == toValue {
					return false
				}
			}
			charTransformationMap[char] = toValue
		}
	}
	return true
}

/*IsAnagram
Given two strings s and t, return true if t is an anagram of s, and false otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
typically using all the original letters exactly once.
Constraints:
- 1 <= s.length, t.length <= 5 * 104
- s and t consist of lowercase English letters.
From https://leetcode.com/problems/valid-anagram/
*/
func IsAnagram(s string, t string) bool {
	charOccurrencesS := make(map[int32]int, 0)
	for _, char := range s {
		charOccurrencesS[char] += 1
	}
	for _, char := range t {
		if value, isExist := charOccurrencesS[char]; isExist {
			if value <= 0 {
				return false
			} else {
				charOccurrencesS[char] -= 1
			}
		} else {
			return false
		}
	}
	for _, value := range charOccurrencesS {
		if value != 0 {
			return false
		}
	}
	return true
}

/*DefangIPaddr
Given a valid (IPv4) IP address, return a defanged version of that IP address.

A defanged IP address replaces every period "." with "[.]".

Constraints:
- The given address is a valid IPv4 address.
*/
func DefangIPaddr(address string) string {
	result := ""
	for _, char := range address {
		if char == '.' {
			result = result + "[.]"
		} else {
			result = result + string(char)
		}
	}
	return result
}

/*FinalValueAfterOperations
There is a programming language with only four operations and one variable X:

++X and X++ increments the value of the variable X by 1.
--X and X-- decrements the value of the variable X by 1.
Initially, the value of X is 0.

Given an array of strings operations containing a list of operations,
return the final value of X after performing all the operations.

Constraints:
- 1 <= operations.length <= 100
- operations[i] will be either "++X", "X++", "--X", or "X--"

From https://leetcode.com/problems/final-value-of-variable-after-performing-operations/
*/
func FinalValueAfterOperations(operations []string) int {
	result := 0
	valueMap := map[uint8]int{
		'-': -1,
		'X': 0,
		'+': 1,
	}
	for _, str := range operations {
		result += valueMap[str[0]]
		result += valueMap[str[2]]
	}
	return result
}

/*MinPartitions
A decimal number is called deci-binary if each of its digits is either 0 or 1 without any leading zeros.
For example, 101 and 1100 are deci-binary, while 112 and 3001 are not.

Given a string n that represents a positive decimal integer, return the minimum number
of positive deci-binary numbers needed so that they sum up to n.

Constraints:
- 1 <= n.length <= 105
- n consists of only digits.
- n does not contain any leading zeros and represents a positive integer.

From https://leetcode.com/problems/partitioning-into-minimum-number-of-deci-binary-numbers/
*/
func MinPartitions(n string) int {
	maxDigit := '0'
	for _, char := range n {
		if char > maxDigit {
			maxDigit = char
		}
	}
	return int(maxDigit - 48)
}
