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

/*MinOperations
You have n boxes. You are given a binary string boxes of length n, where boxes[i] is '0' if
the ith box is empty, and '1' if it contains one ball.
In one operation, you can move one ball from a box to an adjacent box.
Box i is adjacent to box j if abs(i - j) == 1. Note that after doing so, there may be more than one ball in some boxes.
Return an array answer of size n, where answer[i] is the minimum number of operations
needed to move all the balls to the ith box.
Each answer[i] is calculated considering the initial state of the boxes.

Constraints:

- n == boxes.length
- 1 <= n <= 2000
- boxes[i] is either '0' or '1'.
*/
func MinOperations(boxes string) []int {
	result := make([]int, len(boxes))
	index := 0
	for index < len(boxes) {
		val := int(boxes[index] - '0')
		i := index - 1
		for i >= 0 {
			result[i] = result[i] + (val * (index - i))
			i--
		}

		j := index + 1
		for j < len(boxes) {
			result[j] = result[j] + (val * (j - index))
			j++
		}
		index += 1
	}
	return result
}

/*ExecuteInstructions
There is an n x n grid, with the top-left cell at (0, 0) and the bottom-right cell at (n - 1, n - 1).
You are given the integer n and an integer array startPos where startPos = [startrow, startcol] indicates
that a robot is initially at cell (startrow, startcol).
You are also given a 0-indexed string s of length m where s[i] is the ith instruction for the
robot: 'L' (move left), 'R' (move right), 'U' (move up), and 'D' (move down).
The robot can begin executing from any ith instruction in s. It executes the
instructions one by one towards the end of s but it stops if either of these conditions is met:
The next instruction will move the robot off the grid.
There are no more instructions left to execute.
Return an array answer of length m where answer[i] is the number of instructions
the robot can execute if the robot begins executing from the ith instruction in s.
*/
func ExecuteInstructions(n int, startPos []int, s string) []int {
	type coord struct {
		X int
		Y int
	}
	charCoordMap := map[uint8]coord{
		'R': {X: 1, Y: 0},
		'D': {X: 0, Y: 1},
		'L': {X: -1, Y: 0},
		'U': {X: 0, Y: -1},
	}
	cumulativeCords := make([]coord, len(s))

	cumulativeCords[len(s)-1].X = charCoordMap[s[len(s)-1]].X + startPos[0]
	cumulativeCords[len(s)-1].Y = charCoordMap[s[len(s)-1]].Y + startPos[1]
	fmt.Println(s)
	fmt.Printf("Start Pos X: %d , Y: %d \n", startPos[0], startPos[1])
	fmt.Printf("%d'th AA Coords X: %d , Y: %d\n", len(s)-1, cumulativeCords[len(s)-1].X, cumulativeCords[len(s)-1].Y)
	for i := len(s) - 2; i >= 0; i-- {
		move := charCoordMap[s[i]]
		cumulativeCords[i].X = cumulativeCords[i+1].X + move.X
		cumulativeCords[i].Y = cumulativeCords[i+1].Y + move.Y
		fmt.Printf("%d'th Coords X: %d , Y: %d\n", i, cumulativeCords[i].X, cumulativeCords[i].Y)
	}
	result := make([]int, len(cumulativeCords))
	for i := 0; i < len(cumulativeCords); i++ {
		if cumulativeCords[i].X >= 0 && cumulativeCords[i].X < n &&
			cumulativeCords[i].Y >= 0 && cumulativeCords[i].Y < n {
			result[i] = len(cumulativeCords) - i
			fmt.Printf("%d'th Result %d \n", i, result[i])
		}
	}
	return nil
}

/*MostWordsFound
A sentence is a list of words that are separated by a single space with no leading or trailing spaces.

You are given an array of strings sentences, where each sentences[i] represents a single sentence.

Return the maximum number of words that appear in a single sentence.

Constraints:
	- 1 <= sentences.length <= 100
	- 1 <= sentences[i].length <= 100
	- sentences[i] consists only of lowercase English letters and ' ' only.
	- sentences[i] does not have leading or trailing spaces.
	- All the words in sentences[i] are separated by a single space.
*/
func MostWordsFound(sentences []string) int {
	mostWordsFound := 0
	for _, sentence := range sentences {
		if a := len(strings.Split(sentence, " ")); a > mostWordsFound {
			mostWordsFound = a
		}
	}
	return mostWordsFound
}

/*NumJewelsInStones
You're given strings jewels representing the types of stones
that are jewels, and stones representing the stones you have.
Each character in stones is a type of stone you have. You want to know how many
of the stones you have are also jewels.
Letters are case sensitive, so "a" is considered a different type of stone from "A".
Constraints:
	1 <= jewels.length, stones.length <= 50
	jewels and stones consist of only English letters.
	All the characters of jewels are unique.
*/
func NumJewelsInStones(jewels string, stones string) int {
	jewelsCount := 0
	jewelsMap := make(map[int32]bool, len(jewels))
	for _, char := range jewels {
		jewelsMap[char] = true
	}
	for _, char := range stones {
		if jewelsMap[char] {
			jewelsCount += 1
		}
	}
	return jewelsCount
}
func GoalParserInterpretation(command string) string {
	i := 0
	goalParsedResult := ""
	for i < len(command) {
		if command[i] == 'G' {
			goalParsedResult += "G"
			i += 1
		} else if command[i] == '(' && (i+1 < len(command) && command[i+1] == ')') {
			goalParsedResult += "o"
			i += 2
		} else if command[i] == '(' && (i+3 < len(command)) && command[i+1] == 'a' && command[i+2] == 'l' && command[i+3] == ')' {
			goalParsedResult += "al"
			i += 4
		}
	}
	return goalParsedResult
}

/*RestoreString
You are given a string s and an integer array indices
of the same length. The string s will be shuffled such that the character at the ith position moves to indices[i] in the shuffled string.

Return the shuffled string.

Constraints:

	s.length == indices.length == n
	1 <= n <= 100
	s consists of only lowercase English letters.
	0 <= indices[i] < n
	All values of indices are unique.
*/
func RestoreString(s string, indices []int) string {
	a := []uint8(s)
	for i := 0; i < len(s); i++ {
		a[indices[i]] = s[i]
	}
	return string(a)
}

/*CellsInRange
A cell (r, c) of an excel sheet is represented as a string
"<col><row>" where:
<col> denotes the column number c of the cell. It is represented by
alphabetical letters. For example, the 1st column is denoted by
'A', the 2nd by 'B', the 3rd by 'C', and so on.
<row> is the row number r of the cell. The rth row is represented by
the integer r. You are given a string s in the format
"<col1><row1>:<col2><row2>", where <col1> represents the
column c1, <row1> represents the row r1, <col2> represents the
column c2, and <row2> represents the row r2, such that
r1 <= r2 and c1 <= c2.
Return the list of cells (x, y) such that r1 <= x <= r2 and
c1 <= y <= c2. The cells should be represented as strings in the
format mentioned above and be sorted in non-decreasing order first
by columns and then by rows.

Constraints:

	s.length == 5
	'A' <= s[0] <= s[3] <= 'Z'
	'1' <= s[1] <= s[4] <= '9'
	s consists of uppercase English letters, digits and ':'.

*/
func CellsInRange(s string) []string {
	cells := make([]string, 0)
	// (x,y) ->  r1 <= x <= r2
	for j := s[0]; j <= s[3]; j++ {
		for i := int(s[1] - 48); i <= int(s[4]-48); i++ {
			cells = append(cells, fmt.Sprintf("%c%d", j, i))
		}
	}
	return cells
}

/*ReverseWords
Given a string s, reverse the order of characters in
each word within a sentence while still preserving
whitespace and initial word order.

Constraints:

	1 <= s.length <= 5 * 104
	s contains printable ASCII characters.
	s does not contain any leading or trailing spaces.
	There is at least one word in s.
	All the words in s are separated by a single space.
*/
func ReverseWords(s string) string {
	words := strings.Split(s, " ")
	result := make([]string, 0)
	for _, word := range words {
		wordChars := []int32(word)
		j := len(wordChars) - 1
		i := 0
		for j > i {
			wordChars[j] = wordChars[j] + wordChars[i]
			wordChars[i] = wordChars[j] - wordChars[i]
			wordChars[j] = wordChars[j] - wordChars[i]
			i += 1
			j -= 1
		}
		result = append(result, string(wordChars))
	}
	return strings.Join(result, " ")
}
