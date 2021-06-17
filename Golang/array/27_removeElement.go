/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00% 的用户
内存消耗：2.1 MB, 在所有 Go 提交中击败了55.94% 的用户


TODO 参看halfrost
要求:
给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。

不要使用额外的数组空间，你必须仅使用 O(1) 额外空间并 原地 修改输入数组。

元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

*/
package main

import "fmt"

func main() {
	nums := []int{3, 2, 2, 3}
	val := 3
	ret := removeElement(nums, val)
	fmt.Println("The result is", ret)
	fmt.Println("nums", nums)
}

func removeElement(nums []int, val int) int {
	latest, finder := 0, 0
	// find first target elemnt
	for ; finder < len(nums); finder++ {
		if nums[finder] == val {
			latest = finder
			break
		}

	}
	if finder == len(nums) {
		return len(nums)
	}
	for ; finder < len(nums); finder++ {
		if nums[finder] != val {
			nums[latest] = nums[finder]
			latest++
		}
	}
	return latest

}
