package leetcode

import (
	"testing"
)

type Input struct {
	times [][]int
	n     int
	k     int
}

type TestCase struct {
	input  Input
	output int
}

// 输入：times = [[2,1,1],[2,3,1],[3,4,1]], n = 4, k = 2
// 输出：2

// 示例 2：

// 输入：times = [[1,2,1]], n = 2, k = 1
// 输出：1

// 示例 3：

// 输入：times = [[1,2,1]], n = 2, k = 2
// 输出：-1

func Test_743(t *testing.T) {
	test_cases := []TestCase{
		{Input{[][]int{[]int{2, 1, 1}, []int{2, 3, 1}, []int{3, 4, 1}}, 4, 2}, 2},
		{Input{[][]int{[]int{1, 2, 1}}, 2, 1}, 1},
		{Input{[][]int{[]int{1, 2, 1}}, 2, 2}, -1},
		{Input{[][]int{[]int{1, 2, 1}, []int{2, 3, 2}, []int{1, 3, 4}}, 3, 1}, 3},
		{Input{[][]int{[]int{1, 2, 1}, []int{2, 1, 3}}, 2, 2}, 3},
	}
	for _, test_case := range test_cases {
		input := test_case.input
		if res := networkDelayTime(input.times, input.n, input.k); res == test_case.output {
			t.Log("通过")
		} else {
			t.Error("未通过")
		}
	}
}
