/*
Score:
执行用时：116 ms, 在所有 Go 提交中击败了50.30% 的用户
内存消耗：25.9 MB, 在所有 Go 提交中击败了12.10% 的用户
halfrost 还不如我,哈哈


Points:
1. 存在负数
2. 不知道是否存在重复的

Thoughts:
1. 如果排序, 那么时间复杂度为nlog(n)
2. O(0)也就是只许遍历一遍数组
3. halfrost的做法是先将数组字典化, 再从1开始匹配
4. 我的做法是一边匹配一边字典化, 运气好可以(比如有升序部分越多), 时间越少



给你一个未排序的整数数组 nums ，请你找出其中没有出现的最小的正整数。
请你实现时间复杂度为 O(n) 并且只使用常数级别额外空间的解决方案。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/first-missing-positive
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package leetcode

// import "fmt"

func firstMissingPositive(nums []int) int {
	count, search_map := 1, map[int]int{}
	for _, num := range nums {
		if num == count {
			count++
			for {
				if _, exist := search_map[count]; exist {
					count++
				} else {
					break
				}
			}
		} else {
			search_map[num] = 0
		}
	}

	return count
}
