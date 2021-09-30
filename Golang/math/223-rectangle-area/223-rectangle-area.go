/*
Score:
第一例
执行用时：8 ms, 在所有 Go 提交中击败了97.59% 的用户
内存消耗：6.2 MB, 在所有 Go 提交中击败了69.88% 的用户
第二例
执行用时：4 ms, 在所有 Go 提交中击败了100.00% 的用户
内存消耗：6.2 MB, 在所有 Go 提交中击败了69.88% 的用户

Points:
1.

Thoughts:
1. 需要考虑到面积为0
2. 需要考虑重叠
3. 需要考虑包含
4. 仅仅上下镜像转换， 不一定能让第一个矩形靠上

给你 二维 平面上两个 由直线构成的 矩形，请你计算并返回两个矩形覆盖的总面积。

每个矩形由其 左下 顶点和 右上 顶点坐标表示：

    第一个矩形由其左下顶点 (ax1, ay1) 和右上顶点 (ax2, ay2) 定义。
    第二个矩形由其左下顶点 (bx1, by1) 和右上顶点 (bx2, by2) 定义。


https://leetcode-cn.com/problems/rectangle-area/


*/
package leetcode

// import "fmt"

func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	// 假设矩形a靠左， 矩形b靠右
	if ax1 > bx1 {
		return computeArea(bx1, by1, bx2, by2, ax1, ay1, ax2, ay2)
	}
	// 假设矩形a靠上， 矩形b靠下
	if ay2 < by2 {
		if ay1 > by1 {
			// y轴包围
			by1 -= (by2 - ay2)
			by2 = ay2
			return computeArea(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2)
		} else {
			return computeArea(ax1, -ay2, ax2, -ay1, bx1, -by2, bx2, -by1)
		}

		// 沿x轴进行镜像对称变换, 注意1与2组合

	}
	ret := (ay1-ay2)*(ax1-ax2) + (bx1-bx2)*(by1-by2)
	// 排除0面积
	if ay1 == ay2 || ax1 == ax2 || bx1 == bx2 || by1 == by2 {
		return ret
	}
	//分辨重叠
	if ax2 > bx1 && ay1 < by2 {
		// 分辨包含
		w, h := 0, 0
		// x轴是否包含
		if ax2 > bx2 {
			w = bx2 - bx1
			// fmt.Println("x轴包含", w)
		} else {
			w = ax2 - bx1
			// fmt.Println("x轴不包含", w)
		}
		// y轴是否包含
		if ay1 < by1 {
			h = by2 - by1
			// fmt.Println("y轴包含", h)
		} else {
			h = by2 - ay1
			// fmt.Println("y轴不包含", h)
		}
		ret -= w * h
	}

	return ret

}

func computeArea1(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {
	if ax2 < bx1 || bx2 < ax1 || ay2 < by1 || by2 < ay1 {
		// 无重叠 充分必要条件
		return (ay1-ay2)*(ax1-ax2) + (bx1-bx2)*(by1-by2)
	}
	// 确定四个极限点
	left_x, right_x, top_y, bottom_y := 0, 0, 0, 0
	if ax1 < bx1 {
		left_x = bx1
	} else {
		left_x = ax1
	}
	if ax2 < bx2 {
		right_x = ax2
	} else {
		right_x = bx2
	}
	if ay2 < by2 {
		top_y = ay2
	} else {
		top_y = by2
	}
	if ay1 < by1 {
		bottom_y = by1
	} else {
		bottom_y = ay1
	}

	return (ay1-ay2)*(ax1-ax2) + (bx1-bx2)*(by1-by2) - (right_x-left_x)*(top_y-bottom_y)
}
