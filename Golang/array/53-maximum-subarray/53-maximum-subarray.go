/*
Score:
执行用时：4 ms, 在所有 Go 提交中击败了97.05% 的用户
内存消耗：3.2 MB, 在所有 Go 提交中击败了34.42% 的用户

Points:
1.

Thoughts:
1. 以迭代的思维思考， 对当前遍历的一个数而言， 是否将它加入到当前的数列和中还是舍弃之前的数列重新加
2. 遇到负数就可以结算了, 不对[1,2,3,-5,6]，应该是试探性地结算
3. 应该是和一旦小于等于0， 就可以把之前的拿来结算了


给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

*/
package leetcode

// import "fmt"

func maxSubArray(nums []int) int {
	res, tmp := nums[0], nums[0]
	for _, num := range nums[1:] {
		if tmp < 0 {
			tmp = num
			if tmp > res {
				res = tmp
			}
			continue
		}
		if num < 0 {
			if tmp > res {
				res = tmp
			}
		}
		tmp += num

	}

	if nums[len(nums)-1] >= 0 {
		if tmp > res {
			res = tmp
		}
	}
	return res
}
