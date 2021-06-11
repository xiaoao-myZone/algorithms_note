/*
执行用时：28 ms, 在所有 Go 提交中击败了9.11% 的用户
内存消耗：4.5 MB, 在所有 Go 提交中击败了69.27% 的用户
哭了
halfrost
执行用时：8 ms, 在所有 Go 提交中击败了87.43% 的用户
内存消耗：4.5 MB, 在所有 Go 提交中击败了99.96% 的用户

将nums直接传递给函数,只拷贝了slice的三个参数以及内部数组的地址首地址
打印slice的地址其实是内部数组的作为其首部的索引地址

要求:
给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。

不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成


网站最后的判断依据是nums[:ret]
*/
package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 4}
	ret := removeDuplicates_1(nums)
	fmt.Println("The result is ", ret, nums)
	fmt.Printf("%p\n", &nums)
}

func removeDuplicates(nums []int) int {
	// fmt.Printf("%p\n", &nums)
	length, val := len(nums), 0
	if length == 0 {
		return 0
	} else {
		val = nums[0]
	}

	for i := 1; i < length; {
		if val == nums[i] {
			nums = append(nums[:i], nums[i+1:]...)
			length--
		} else {
			val = nums[i]
			i++
		}

	}
	// fmt.Println(nums)
	// fmt.Printf("%p\n", &nums)
	return length
}

func removeDuplicates_1(nums []int) int {
	//妙
	if len(nums) == 0 {
		return 0
	}
	last, finder := 0, 0
	for last < len(nums)-1 {
		for nums[finder] == nums[last] {
			finder++
			if finder == len(nums) {
				return last + 1
			}
		}
		nums[last+1] = nums[finder]
		last++
	}
	return last + 1
}
