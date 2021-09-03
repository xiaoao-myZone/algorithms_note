package leetcode

import (
	"testing"
)

type Input struct {
	haystack string
	needle   string
}

type TestCase struct {
	input  Input
	output int
}

func Test_0_template(t *testing.T) {
	test_cases := []TestCase{
		{Input{"hello", "needle"}, 1},
		{Input{"aaaaa", "bba"}, -1},
		{Input{"", ""}, 0},
	}
	for _, test_case := range test_cases {
		if res := strStr(test_case.input.haystack, test_case.input.needle); res == test_case.output {
			t.Log("通过")
		} else {
			t.Error("未通过", res, "/", test_case.output)
		}
	}
}
