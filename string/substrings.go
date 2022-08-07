package string

import "fmt"

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
		fmt.Printf("Start: %d, Char: %c, Finish: %d, Index: %d \n", longestSubstr.Start, s[index-1], longestSubstr.Finish, index)
	}
	// that means string is ful of not repeated characters.
	//if longestSubstr.Finish != len(s) {
	//	longestSubstr.Finish = len(s) - 1
	//	longestSubstr.Length = longestSubstr.Finish - longestSubstr.Start
	//}
	return longestSubstr.Finish - longestSubstr.Start
}
