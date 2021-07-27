package leetcode

import (
	"../utils"
	"reflect"
	"testing"
)

type TestCase struct {
	input  [][]int
	output []int
}

func Test_1743(t *testing.T) {
	// [2,1],[3,4],[3,2]
	test_cases := []TestCase{
		{
			[][]int{[]int{5, 4}, []int{2, 1}, []int{5, 6}, []int{3, 4}, []int{3, 2}},
			[]int{6, 5, 4, 3, 2, 1},
		},
	}
	for _, test_case := range test_cases {
		if res := restoreArray1(test_case.input); reflect.DeepEqual(res, test_case.output) ||
			reflect.DeepEqual(utils.IntReverse(res, true), test_case.output) {
			//res == test_case.output {
			t.Log("通过")
		} else {
			t.Error("未通过")
		}
	}
}
