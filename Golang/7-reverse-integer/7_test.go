package leetcode

import (
	"testing"
)

type TestCase struct {
	input  int
	output int
}

func Test_0_template(t *testing.T) {
	test_cases := []TestCase{
		{1, 1},
		{21, 12},
		{-321, -123},
		{870, 78},
		{0, 0},
		// {2147483647},
	}
	for _, test_case := range test_cases {
		if res := reverse(test_case.input); res == test_case.output {
			t.Log("通过")
		} else {
			t.Error("未通过", res, "/", test_case.output)
		}
	}
}
