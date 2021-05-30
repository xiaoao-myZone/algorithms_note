/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00% 的用户
内存消耗：2.7 MB, 在所有 Go 提交中击败了75.71% 的用户

要点:
	1. 找出最接近目标值的答案, 假设只存在唯一解
	2. 长度大于三
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{0, 1, 2}
	target := 3
	ret := threeSumClosest(nums, target)
	fmt.Println("The result is ", ret)
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	ret := nums[0] + nums[1] + nums[2] - target
	if ret == 0 {
		return target
	}
	tmp := 0
	left, mid, right := 0, 1, len(nums)-1
	for ; left < len(nums)-2; left++ {
		if left > 0 && nums[left] == nums[left-1] {
			continue
		}
		for mid, right = left+1, len(nums)-1; right > mid; {
			tmp = nums[left] + nums[mid] + nums[right] - target
			if abs(tmp) < abs(ret) {
				if tmp == 0 {
					return target
				} else {
					ret = tmp
					//fmt.Println(nums[left], nums[mid], nums[right])
				}
			} else {
				if tmp > 0 {
					right--
				} else {
					mid++
				}
			}

		}
	}
	return ret + target
}

func abs(n int) int {
	if n > 0 {
		return n
	} else {
		return -n
	}
}
