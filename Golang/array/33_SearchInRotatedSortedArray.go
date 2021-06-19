/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00% 的用户
内存消耗：2.5 MB, 在所有 Go 提交中击败了100.00% 的用户

Point:
1. 未旋转前为升序
2. 数值各不相同

在搜索断层的过程中不耽误找到target
初始需要分三段,左中右, 肯定有一边是保序的
先搜索保序的那部分,再分隔不保序的
一旦取值范围进入保序区间,就要不必再跑去非保序区间
*/
package main

import "fmt"

func main() {
	nums := []int{5, 1, 3} // {5, 1, 3} //{4, 5, 6, 7, 0, 1, 2} // {1,3}
	target := 5
	ret := search(nums, target)
	fmt.Println("The result is ", ret)
}

func search(nums []int, target int) int {

	left, mid, right := 0, 0, len(nums)-1
	mid = right
	if nums[right] == target { //由于mid = (left + right) / 2 偏左, 无法遍历到初始数组的右边界
		return right
	}
	for left < right {
		//fmt.Println(left, mid, right)
		mid = (left + right) / 2 // 偏左, 搜索区间[left, right)
		//fmt.Println(left, mid, right)
		if nums[mid] == target {
			return mid
		}
		if nums[left] <= nums[mid] { //前半段保序, 虽然数值各不相同, 但是left和mid可能相等
			if nums[left] <= target && target < nums[mid] { //在保序区间, 继续缩小范围
				right = mid
			} else {
				left = mid
				left++
			}
		} else { // nums[mid]<nums[right] 后半段保序
			//mid++ 看似缩小了搜索范围(其实只有长度1), 但是如果失败, 会让right增大, 夸大了下一次搜索范围
			// 在{5, 1, 3} 5 这个例子中会出现死循环
			if nums[mid] <= target && target < nums[right] { //在保序区间, 继续缩小范围
				left = mid
			} else {
				right = mid
			}
			//fmt.Println(right, left, mid)
		}

	}
	return -1
}
