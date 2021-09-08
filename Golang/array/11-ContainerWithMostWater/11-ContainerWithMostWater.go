/*
Score:
执行用时：76 ms, 在所有 Go 提交中击败了93.74% 的用户
内存消耗：8.6 MB, 在所有 Go 提交中击败了52.09% 的用户

Points:
1.

Thoughts:
1. 记忆中是双指针
2. 从两边出发
3. 每次移动较短的一边(一样该怎么办? 应该不需要担心, 因为只有两个都移动到比以前高的地方才能大于这个值)



给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0) 。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

说明：你不能倾斜容器。



示例 1：

输入：[1,8,6,2,5,4,8,3,7]
输出：49
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/container-with-most-water
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
package leetcode

// import "fmt"

func maxArea(height []int) int {
	l_index, r_index := 0, len(height)-1
	ret, tmp := 0, 0
	for l_index != r_index {
		if height[l_index] > height[r_index] {
			tmp = (r_index - l_index) * height[r_index]
			r_index -= 1
		} else {
			tmp = (r_index - l_index) * height[l_index]
			l_index += 1
		}
		if tmp > ret {
			ret = tmp
		}
	}
	return ret
}
