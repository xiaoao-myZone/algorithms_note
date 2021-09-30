package leetcode

import (
	"testing"
)

type Input struct {
	dividend int
	divisor  int
}

type TestCase struct {
	input  Input
	output int
}

func Test_0_template(t *testing.T) {
	test_cases := []TestCase{
		{Input{10, 3}, 3},
		{Input{7, -3}, -2},
		{Input{-2147483648, -1}, 2147483647},
		{Input{2147483647, -1}, -2147483647},
	}
	for _, test_case := range test_cases {
		input := test_case.input
		if res := divide1(input.dividend, input.divisor); res == test_case.output {
			t.Log("通过")
		} else {
			t.Error("未通过", res, "/", test_case.output)
		}
	}
}
