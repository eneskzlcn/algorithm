package string_algorithms

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
