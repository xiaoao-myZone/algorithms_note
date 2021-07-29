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
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6},
		{[]int{4, 2, 0, 3, 2, 5}, 9},
	}
	for _, test_case := range test_cases {
		if res := trap(test_case.input); res == test_case.output {
			t.Log("通过")
		} else {
			t.Error("未通过")
		}
	}
}
