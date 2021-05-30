/*
执行用时：100 ms, 在所有 Go 提交中击败了69.52% 的用户
内存消耗：8.3 MB, 在所有 Go 提交中击败了62.82% 的用户
halfrost的解法与我基本一样

1. 记忆中是双指针
2. 从两边出发
3. 每次移动较短的一边(一样该怎么办? 应该不需要担心, 因为只有两个都移动到比以前高的地方才能大于这个值)
*/
package main

import "fmt"

func maxArea(height []int) int {
	l_index, r_index := 0, len(height)-1
	ret, tmp := 0, 0
	for l_index != r_index {
		if height[l_index] > height[r_index] {
			tmp = (r_index - l_index) * height[r_index]
			r_index -= 1
		} else {
			tmp = (r_index - l_index) * height[l_index]
			l_index += 1
		}
		if tmp > ret {
			ret = tmp
		}
	}
	return ret
}

func main() {
	nums := []int{1, 2, 3}
	ret := maxArea(nums)
	fmt.Println("return: ", ret)
}
