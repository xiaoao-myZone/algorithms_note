/*
Score:
执行用时：8 ms, 在所有 Go 提交中击败了98.81% 的用户
内存消耗：5.2 MB, 在所有 Go 提交中击败了72.42% 的用户

Points:
1. 须考虑空数组， 但不会都是空
2. 数组中有相同的值

Thoughts:
1.


给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

示例 1：

输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
package leetcode

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// i, j 用来计数， mid标记中点， current用来存当前数据， tmp用来存上一次数据, 以防出现偶数长度个数列的中位数
	i, j, mid, current, tmp := 0, 0, (len(nums1)+len(nums2))/2+1, 0, 0
	for i+j < mid {
		tmp = current
		// 这总写法虽然增加了一次判断， 但是可以减少代码的缩进深度
		if j < len(nums2) && i < len(nums1) {
			if nums1[i] < nums2[j] {
				current = nums1[i]
				i++
			} else {
				current = nums2[j]
				j++
			}
		} else if i < len(nums1) {
			current = nums1[i]
			i++
		} else {
			current = nums2[j]
			j++
		}
	}
	if (len(nums1)+len(nums2))%2 == 0 {
		return float64(tmp+current) / 2
	} else {
		return float64(current)
	}

}
