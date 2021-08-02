package leetcode

import (
	"testing"
)

type TestCase struct {
	input  []int
	output int
}

func Test_48(t *testing.T) {
	test_cases := []TestCase{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{1}, 1},
		{[]int{0}, 0},
		{[]int{-1}, -1},
		{[]int{-100000}, -100000},
		{[]int{5, 4, -1, 7, 8}, 23},
		{[]int{2, 3, 0}, 5},
	}
	for _, test_case := range test_cases {
		if res := maxSubArray(test_case.input); res == test_case.output {
			t.Log("通过")
		} else {
			t.Error("未通过")
		}
	}
}
