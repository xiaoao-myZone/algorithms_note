package leetcode

import (
	"testing"
)

type TestCase struct {
	input  int
	output string
}

func Test_0_template(t *testing.T) {
	test_cases := []TestCase{
		{3, "III"},
		{4, "IV"},
		{9, "IX"},
		{58, "LVIII"},
		{1994, "MCMXCIV"},
	}
	for _, test_case := range test_cases {
		if res := intToRoman(test_case.input); res == test_case.output {
			t.Log("通过")
		} else {
			t.Error("未通过", res, "/", test_case.output)
		}
	}
}
