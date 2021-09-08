package leetcode

import (
	"testing"
)

type Input struct {
	nums   []int
	target int
}
type TestCase struct {
	input  Input
	output int
}

func Test_0_template(t *testing.T) {
	test_cases := []TestCase{
		{Input{[]int{-1, 2, 1, -4}, 1}, 2},
	}
	for _, test_case := range test_cases {
		input := test_case.input
		if res := threeSumClosest(input.nums, input.target); res == test_case.output {
			t.Log("通过")
		} else {
			t.Error("未通过", res, "/", test_case.output)
		}
	}
}
