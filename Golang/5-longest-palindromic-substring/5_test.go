package leetcode

import (
	"testing"
)

type TestCase struct {
	input  string
	output string
}

func Test_0_template(t *testing.T) {
	test_cases := []TestCase{
		{"babad", "aba"},
		{"cbbd", "bb"},
		{"a", "a"},
		{"ac", "a"},
	}
	for _, test_case := range test_cases {
		if res := longestPalindrome(test_case.input); res == test_case.output {
			t.Log("通过")
		} else {
			t.Error("未通过", res, "/", test_case.output)
		}
	}
}
