package leetcode

import (
	"testing"
)

type TestCase struct {
	input  string
	output bool
}

func Test_551(t *testing.T) {
	test_cases := []TestCase{
		{"PPALLP", true},
	}
	for _, test_case := range test_cases {
		if res := checkRecord(test_case.input); res == test_case.output {
			t.Log("通过")
		} else {
			t.Error("未通过")
		}
	}
}
