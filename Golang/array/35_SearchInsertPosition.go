/*
执行用时：4 ms, 在所有 Go 提交中击败了91.20% 的用户
内存消耗：3 MB, 在所有 Go 提交中击败了95.95% 的用户

Points:
1. 要么找到给定数在有序数组中的位置, 要么给出插入的位置
2. 数组长度不为0
3. 无重复元素
*/
package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 6}
	target := 4
	ret := searchInsert(nums, target)
	fmt.Println("The result is ", ret)
}

func searchInsert(nums []int, target int) int {
	left, mid, right := 0, 0, len(nums)-1
	if nums[right] == target {
		return right
	}
	if nums[right] < target {
		return right + 1
	}
	for left < right {
		mid = (left + right) / 2
		if nums[mid] < target {
			left = mid
			left++
		} else if nums[mid] > target {
			right = mid
		} else {
			return mid
		}
	}
	// left == right
	if nums[left] < target {
		return left + 1
	} else {
		return left
	}

}
