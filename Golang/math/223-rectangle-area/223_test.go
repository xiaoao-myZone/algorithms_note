package leetcode

import (
	"testing"
)

type TestCase struct {
	input  []int
	output int
}

func Test_0_template(t *testing.T) {
	test_cases := []TestCase{
		{[]int{-3, 0, 3, 4, 0, -1, 9, 2}, 45},
		{[]int{0, 0, 0, 0, -1, -1, 1, 1}, 4},
		{[]int{-2, -2, 2, 2, -1, -1, 1, 1}, 16},
		{[]int{-2, -2, 2, 2, -2, 2, 2, 4}, 24},
		{[]int{-2, -2, 2, 2, 1, -3, 3, 3}, 24},
	}
	for _, test_case := range test_cases {
		i := test_case.input
		if res := computeArea1(i[0], i[1], i[2], i[3], i[4], i[5], i[6], i[7]); res == test_case.output {
			t.Log("通过")
		} else {
			t.Error("未通过", res, "/", test_case.output)
		}
	}
}
