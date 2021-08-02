package leetcode

import (
	"testing"
)

type TestCase struct {
	input  []int
	output int
}

func Test_41(t *testing.T) {
	test_cases := []TestCase{
		{[]int{2, 0, 1}, 3},
		{[]int{3, 4, -1, 1}, 2},
		{[]int{7, 8, 9, 11, 12}, 1},
	}
	for _, test_case := range test_cases {
		if res := firstMissingPositive(test_case.input); res == test_case.output {
			t.Log("通过")
		} else {
			t.Error("未通过")
		}
	}
}
