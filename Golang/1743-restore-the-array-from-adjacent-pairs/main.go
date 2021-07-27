/*
Score:
第二例
执行用时：220 ms, 在所有 Go 提交中击败了31.88% 的用户
内存消耗：29.1 MB, 在所有 Go 提交中击败了81.16% 的用户

执行用时：216 ms, 在所有 Go 提交中击败了34.78% 的用户
内存消耗：31.6 MB, 在所有 Go 提交中击败了62.32% 的用户
# TODO

Points:
1.

Thoughts:
1. 什么情况下会存在多种解答? 任何一种情况都存在至少两个解, 正续和逆序
2. 除了首尾, 每个数都会在两个数对中出现



存在一个由 n 个不同元素组成的整数数组 nums ，但你已经记不清具体内容。好在你还记得 nums 中的每一对相邻元素。

给你一个二维整数数组 adjacentPairs ，大小为 n - 1 ，其中每个 adjacentPairs[i] = [ui, vi] 表示元素 ui 和 vi 在 nums 中相邻。

题目数据保证所有由元素 nums[i] 和 nums[i+1] 组成的相邻元素对都存在于 adjacentPairs 中，存在形式可能是 [nums[i], nums[i+1]] ，也可能是 [nums[i+1], nums[i]] 。这些相邻元素对可以 按任意顺序 出现。

返回 原始数组 nums 。如果存在多种解答，返回 其中任意一个 即可。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/restore-the-array-from-adjacent-pairs

*/
package leetcode

import (
	"fmt"
)

func restoreArray(adjacentPairs [][]int) []int {
	res, search_map, count := make([]int, len(adjacentPairs)+1), map[int][]int{}, 0
	for _, pairs := range adjacentPairs {
		left := pairs[0]
		right := pairs[1]
		if val, ok := search_map[left]; ok {
			search_map[left] = append(val, right)
		} else {
			search_map[left] = []int{right}
		}
		if val, ok := search_map[right]; ok {
			search_map[right] = append(val, left)
		} else {
			search_map[right] = []int{left}
		}

	}
	//fmt.Println(search_map) //空间复杂度2n
	//找到端点
	for k, v := range search_map {
		if len(v) == 1 {
			res[0] = k
			res[1] = v[0]
			count = 2
			// res = append(res, k, v[0])
			// delete(search_map, k)
			break
		}
	}
	// 按图索骥
	for i, j := res[0], res[1]; ; count++ {
		val := search_map[j]
		if len(val) == 1 {
			break
		}
		if val[0] == i {
			res[count] = val[1]
			i, j = j, val[1]
		} else {
			res[count] = val[0]
			i, j = j, val[0]
		}

	}
	return res

}

func restoreArray1(adjacentPairs [][]int) []int {
	// 从出度,入度以及有没有其他数对的元素指向它分析
	res, search_map := []int{}, map[int][2][2]int{} //三元组, 邻接的值, 出度, 入度
	for _, pairs := range adjacentPairs {
		left := pairs[0]
		right := pairs[1]
		// 所有的映射都只需要保证有出度
		if val, ok := search_map[left]; ok {
			// 它已经指着别人了
			// 往前延伸到第三数的入度
			if search_map[val[0]][1] == 1 {
			}
		} else {
			if val, ok := search_map[right]; ok {
			} else {

			}
		}

	}
	fmt.Println(search_map)

	return res
}
