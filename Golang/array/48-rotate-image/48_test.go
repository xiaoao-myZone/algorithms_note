package leetcode

import (
	"reflect"
	"testing"
)

type TestCase struct {
	input  [][]int
	output [][]int
}

func Test_48(t *testing.T) {
	test_cases := []TestCase{
		{
			[][]int{[]int{1, 2}, []int{4, 3}},
			[][]int{[]int{4, 1}, []int{3, 2}},
		},
		{
			[][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}},
			[][]int{[]int{7, 4, 1}, []int{8, 5, 2}, []int{9, 6, 3}},
		},
		{
			[][]int{[]int{5, 1, 9, 11}, []int{2, 4, 8, 10}, []int{13, 3, 6, 7}, []int{15, 14, 12, 16}},
			[][]int{[]int{15, 13, 2, 5}, []int{14, 3, 4, 1}, []int{12, 6, 8, 9}, []int{16, 7, 10, 11}},
		},
	}
	for _, test_case := range test_cases {
		if rotate(test_case.input); reflect.DeepEqual(test_case.input, test_case.output) {
			t.Log("通过")
		} else {
			t.Error("未通过")
		}
	}
}
