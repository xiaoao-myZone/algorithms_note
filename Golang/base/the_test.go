package base

import (
	"testing"
)

type Input struct {
	array  []int
	target int
}

type TestCase struct {
	input  Input
	output int
}

func Test(t *testing.T) {
	test_cases := []TestCase{
		{Input{[]int{1, 2, 3, 5, 6}, 4}, 3},
		{Input{[]int{6, 8, 11, 20, 34}, 15}, 3},
		{Input{[]int{1, 2, 3, 5, 6}, 7}, 5},
		{Input{[]int{1, 2, 3, 5, 6}, 0}, 0},
		{Input{[]int{1, 2, 3, 5, 6}, 2}, -1},
		{Input{[]int{1, 2, 3, 5, 6}, 6}, -1},
		{Input{[]int{1, 2, 3, 5, 6}, 1}, -1},
		{Input{[]int{1}, 0}, 0},
		{Input{[]int{1}, 2}, 1},
	}
	for _, test_case := range test_cases {
		input := test_case.input
		if res := SearchInsertPoint(input.array, input.target); res == test_case.output {
			t.Log("通过")
		} else {
			t.Error("不通过: ", res)
		}
	}
}
