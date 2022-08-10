package dynamic_programming

/*CountBits
Given an integer n, return an array ans of length n + 1 such that for each
i (0 <= i <= n), ans[i] is the number of 1's in the binary representation of i.

Constraints:
- 0 <= n <= 105
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
