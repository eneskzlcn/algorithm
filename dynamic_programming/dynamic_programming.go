package dynamic_programming

/*CountBits
Given an integer n, return an array ans of length n + 1 such that for each
i (0 <= i <= n), ans[i] is the number of 1's in the binary representation of i.

Constraints:
- 0 <= n <= 105

From https://leetcode.com/problems/counting-bits/
*/
func CountBits(n int) []int {
	if n == 0 {
		return []int{0}
	}
	if n == 1 {
		return []int{0, 1}
	}
	calculateBinaryOnesInNumberFn := func(number int, result []int) int {
		if number%2 == 0 { //even
			return result[number/2]
		} else {
			return 1 + result[number/2]
		}
	}
	result := make([]int, n+1)
	result[0] = 0
	result[1] = 1
	for i := 0; i < n+1; i++ {
		result[i] = calculateBinaryOnesInNumberFn(i, result)
	}
	return result
}

/*PascalTriangle algorithm generates a pascal triangle as array representation from given number.
Given an integer numRows, return the first numRows of Pascal's triangle.

In Pascal's triangle, each number is the sum of the two numbers above.

Constraints:

- 1 <= numRows <= 30
From https://leetcode.com/problems/pascals-triangle/
*/
//func PascalTriangle(numRows int) [][]int {
//
//}
//func generateRow(number int, result [][]int) {
//
//}
