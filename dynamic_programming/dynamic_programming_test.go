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
