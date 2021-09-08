package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	input  []int
	output [][]int
}

func Test_0_template(t *testing.T) {
	test_cases := []TestCase{
		{[]int{-1, 0, 1, 2, -1, -4}, [][]int{[]int{-1, -1, 2}, []int{-1, 0, 1}}},
		{[]int{}, [][]int{}},
		{[]int{0}, [][]int{}},
	}
	for _, test_case := range test_cases {
		if res := threeSum1(test_case.input); reflect.DeepEqual(res, test_case.output) {
			t.Log("通过")
		} else {
			t.Error("未通过", res, "/", test_case.output)
		}
	}
}
