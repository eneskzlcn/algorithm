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
	if len(stack) > 0 {
		return false
	}
	return true
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
	aValue, _ := strconv.ParseInt(a, 2, 128)
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

const CharA int = 65

func ConvertToTitle(columnNumber int) string {
	value := columnNumber
	result := ""
	for value > 0 {
		bit := (value - 1) % 26 // 1
		value = value - bit
		value = (value - 1) / 26
		result = fmt.Sprintf("%c", int32(bit+CharA)) + result
	}
	return result
}
