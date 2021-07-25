package leetcode

import (
	"testing"
)

type TestCase struct {
	input  string
	output string
}

func Test_1736(t *testing.T) {
	test_cases := []TestCase{
		{"?4:12", "14:12"},
		{"??:??", "23:59"},
		{"2?:3?", "23:39"},
		{"?3:?4", "23:54"},
	}
	for _, test_case := range test_cases {
		if res := maximumTime(test_case.input); res == test_case.output {
			t.Log("通过")
		} else {
			t.Error("未通过")
		}
	}
}
