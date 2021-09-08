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
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
		{[]int{4, 3, 2, 1, 4}, 16},
		{[]int{1, 2, 1}, 2},
	}
	for _, test_case := range test_cases {
		if res := maxArea(test_case.input); res == test_case.output {
			t.Log("通过")
		} else {
			t.Error("未通过", res, "/", test_case.output)
		}
	}
}
