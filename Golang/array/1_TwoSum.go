/*
case-0
执行用时：4 ms, 在所有 Go 提交中击败了90.55% 的用户
内存消耗：3 MB, 在所有 Go 提交中击败了99.99% 的用户
case-2
执行用时：4 ms, 在所有 Go 提交中击败了90.60% 的用户
内存消耗：3 MB, 在所有 Go 提交中击败了99.99% 的用户
*/
package main

import "fmt"

func main() {
	nums := []int{3, 2, 4} //{2, 7, 11, 15}
	var target int = 6
	fmt.Println(nums, target)
	ret := twoSum_2(nums, target)
	fmt.Println(ret)
	// {3, 2, 4}, 6这个输入很容易出错

}

func twoSum(nums []int, target int) []int {

	for i := range nums {
		another := target - nums[i]
		a_nums := nums[i+1:]
		for j := range a_nums {
			if another == a_nums[j] {
				return []int{i, j + i + 1}
			}
		}
	}
	return []int{}
}

func twoSum_1(nums []int, target int) []int {
	// 不能解决有多个相同的值的问题
	m := map[int]int{}
	for i := range nums {
		m[nums[i]] = i
	}
	for k, v := range m {
		if index, ok := m[target-k]; ok {
			return []int{v, index}
		}
	}
	return nil
}

func twoSum_2(nums []int, target int) []int {
	m := map[int]int{}
	for k, v := range nums {
		if i, ok := m[target-v]; ok {
			return []int{i, k}
		}
		m[v] = k
	}
	return nil
}

/*
Conclusion:
	1. twoSum_1不能解决有多个相同的值得问题
	2. twoSum_2"记忆"尝试失败的kv值
	3. twoSum_2找到的结果中的其中一个是从之前尝试失败的记忆中找出来的
	4. 如果存在相同的值, 并且它们相加不是目标值, 之前的值会被覆盖, 在这个题中是无所谓的
	4. 当结果之一在最后时, 有最长搜索时间, 实现了O(n)
	5. leetcode的测试中, 与原方法相比, 差不了多少

*/
