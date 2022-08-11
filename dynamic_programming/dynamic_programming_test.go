package dynamic_programming_test

import (
	"github.com/eneskzlcn/algorithm/dynamic_programming"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCountBits(t *testing.T) {
	type testCase struct {
		Given    int
		Expected []int
	}
	testCases := []testCase{
		{
			Given:    2,
			Expected: []int{0, 1, 1},
		},
		{
			Given:    5,
			Expected: []int{0, 1, 1, 2, 1, 2},
		},
		{
			Given:    3,
			Expected: []int{0, 1, 1, 2},
		},
	}
	for _, test := range testCases {
		result := dynamic_programming.CountBits(test.Given)
		assert.ElementsMatchf(t, test.Expected, result, "")
	}
}

func TestPascalTriangle(t *testing.T) {
	type testCase struct {
		Given    int
		Expected [][]int
	}
	testCases := []testCase{
		//{
		//	Given:    5,
		//	Expected: [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}},
		//},
		{
			Given:    1,
			Expected: [][]int{{1}},
		},
	}
	for _, test := range testCases {
		result := dynamic_programming.PascalTriangle(test.Given)
		for index, list := range result {
			assert.ElementsMatchf(t, test.Expected[index], list, "")

		}
	}
}
