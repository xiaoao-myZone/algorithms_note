package leetcode

import (
	"testing"
)

type TestCase struct {
	input  []int
	output bool
}

func Test_0_template(t *testing.T) {
	test_cases := []TestCase{
		{[]int{0, 1, 3, 5, 6, 8, 12, 17}, true},
		{[]int{0, 1, 2, 3, 4, 8, 9, 11}, false},
	}
	for _, test_case := range test_cases {
		if res := canCross(test_case.input); res == test_case.output {
			t.Log("通过")
		} else {
			t.Error("未通过", res, "/", test_case.output)
		}
	}
}
