package leetcode

import (
	"testing"
)

type TestCase struct {
	input  [2][]int
	output float64
}

// 示例 1：

// 输入：nums1 = [1,3], nums2 = [2]
// 输出：2.00000
// 解释：合并数组 = [1,2,3] ，中位数 2

// 示例 2：

// 输入：nums1 = [1,2], nums2 = [3,4]
// 输出：2.50000
// 解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5

// 示例 3：

// 输入：nums1 = [0,0], nums2 = [0,0]
// 输出：0.00000

// 示例 4：

// 输入：nums1 = [], nums2 = [1]
// 输出：1.00000

// 示例 5：

// 输入：nums1 = [2], nums2 = []
// 输出：2.00000

func Test_4(t *testing.T) {
	test_cases := []TestCase{
		{[2][]int{[]int{1, 3}, []int{2}}, 2.0000},
		{[2][]int{[]int{1, 2}, []int{3, 4}}, 2.5},
		{[2][]int{[]int{0, 0}, []int{0, 0}}, 0.0},
		{[2][]int{[]int{}, []int{1}}, 1.0},
		{[2][]int{[]int{2}, []int{}}, 2.00},
	}
	for _, test_case := range test_cases {
		input := test_case.input
		if res := findMedianSortedArrays(input[0], input[1]); res == test_case.output {
			t.Log("通过")
		} else {
			t.Error("未通过", res, "/", test_case.output)
		}
	}
}
