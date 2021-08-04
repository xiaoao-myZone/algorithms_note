package leetcode

import (
	"testing"
)

type TestCase struct {
	input  []int
	output int
}

func Test_0_template(t *testing.T) {
	test_cases := []TestCase{
		{[]int{2, 6, 4, 8, 10, 9, 15}, 5},
		{[]int{1, 2, 3, 4}, 0},
		{[]int{1}, 0},
		{[]int{1, 3, 2, 3, 3}, 2},
	}
	for _, test_case := range test_cases {
		if res := findUnsortedSubarray1(test_case.input); res == test_case.output {
			t.Log("通过")
		} else {
			t.Error("未通过", res)
		}
	}
}
