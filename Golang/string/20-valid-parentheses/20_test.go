package leetcode

import (
	"testing"
)

type TestCase struct {
	input  string
	output bool
}

func Test_0_template(t *testing.T) {
	test_cases := []TestCase{
		// {"()", true},
		// {"()[]{}", true},
		// {"(]", false},
		// {"([)]", false},
		// {"{[]}", true},
		// {"(([]){})", true},
		{"({{{{}}}))", false},
	}
	for _, test_case := range test_cases {
		if res := isValid(test_case.input); res == test_case.output {
			t.Log("通过")
		} else {
			t.Error("未通过", res, "/", test_case.output)
		}
	}
}
