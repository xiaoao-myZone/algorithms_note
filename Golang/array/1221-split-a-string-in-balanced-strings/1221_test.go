package leetcode

import (
	"testing"
)

type TestCase struct {
	input  string
	output int
}

func Test_0_template(t *testing.T) {
	test_cases := []TestCase{
		{"RLRRLLRLRL", 4},
		{"RLLLLRRRLR", 3},
		{"LLLLRRRR", 1},
		{"RLRRRLLRLL", 2},
	}
	for _, test_case := range test_cases {
		if res := balancedStringSplit(test_case.input); res == test_case.output {
			t.Log("通过")
		} else {
			t.Error("未通过", res, "/", test_case.output)
		}
	}
}
