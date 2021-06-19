/*
执行用时：8 ms, 在所有 Go 提交中击败了88.18% 的用户
内存消耗：3.9 MB, 在所有 Go 提交中击败了94.86% 的用户

Points:
1. 数组长度可能为0
2.
*/
package main

import "fmt"

func main() {
	nums := []int{1, 1, 2} //{5, 7, 7, 8, 8, 10}
	target := 1
	ret := searchRange(nums, target)
	fmt.Println("The result is ", ret)
}
func searchRange(nums []int, target int) []int {

	left, mid, right := 0, 0, len(nums)-1
	ret := []int{-1, -1}
	if right < 0 {
		return ret
	}
	if nums[right] == target {
		ret[1] = right
	}
	for left < right {
		mid = (left + right) / 2 //偏左
		if nums[mid] > target {
			right = mid
		} else if nums[mid] < target {
			left = mid
			left++ //注意这种不对称性
		} else {
			//下界
			ret[0] = mid
			mid++
			for mid < len(nums) {
				if nums[mid] == target {
					mid++
				} else {
					ret[1] = mid - 1
					break
				}
			}
			//上界
			mid = ret[0] - 1
			ret[0] = 0 //精髓
			for mid >= 0 {
				if nums[mid] == target {
					mid--
				} else {
					ret[0] = mid + 1
					break
				}
			}
			break
		}
	}
	if ret[1] >= 0 && ret[0] < 0 {
		ret[0] = ret[1]
	}
	return ret
}
