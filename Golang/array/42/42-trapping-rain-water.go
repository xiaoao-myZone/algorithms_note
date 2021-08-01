/*
Score:
执行用时：4 ms, 在所有 Go 提交中击败了77.69% 的用户
内存消耗：2.8 MB, 在所有 Go 提交中击败了100.00% 的用户

Points:
1.

Thoughts:
1. 要么用极值的概念, 要么用栈的概念
2. 后者代码更少一点
3. 这题的巧妙之处在于理解水是如何填满水的槽， 栈的思维可以解决从一侧按顺序填满的过程， 而不是天降暴雨
4. 栈中存储的对象至少是一个二元数组
5. 不一定， 栈中的元素在一定是连续的，只用知道初始索引和长度即可
6. 本质是双(三)指针



给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。



示例 1：

输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。

*/
package leetcode

func trap(height []int) int {
	//fmt.Println("************")
	// origin 到 left 之间是一个非升序列
	res, origin, shorter := 0, 0, 0
	//fmt.Println(height[0], "初始高度")
	for right := 1; right < len(height); right++ {
		//fmt.Println("***", height[origin:right+1])
		if height[right] > height[right-1] { // 蓄水池上升过程，需要结算
			//fmt.Println(height[right], "开始蓄水")
			if height[right] > height[origin] {
				shorter = height[origin]
			} else {
				shorter = height[right]
			}
			//fmt.Println("	短边: ", shorter)
			for index := right - 1; index >= 0 && height[index] < shorter; index-- {
				res += shorter - height[index]
				//fmt.Println("		增加: ", shorter-height[index])
				height[index] = shorter // 填平水已经浸没的地方
			}
			if height[origin] == shorter {
				origin = right

			}
		} //else { //下降或者平流过程
		//fmt.Println(height[right], "下降或者平流")
		//}
		//fmt.Println("current res ", res)
	}
	return res
}
