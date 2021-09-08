/*
Score:
执行用时：588 ms, 在所有 Go 提交中击败了5.03% 的用户
内存消耗：26.5 MB, 在所有 Go 提交中击败了5.00% 的用户
哭了

执行用时：32 ms, 在所有 Go 提交中击败了93.90% 的用户
内存消耗：7.5 MB, 在所有 Go 提交中击败了43.84% 的用户

Points:
1. 返回一个数组中三个不同值且其和等于0的所有解

Thoughts:
	1. 貌似最多只能做到O(n2)
	2. 可以两两组合和, 用0减去, 在列表中搜索该值
	3. 由于这个数组有重复的元素, map的值是一个分片,分片由index组成
	4. 在搜索时需要把内外循环中的index从map对应的值中移除, 并且在循环结束后加上
	5. 为了防止出现重复的结果, 需要在搜索前, 建立好map
	6. 搜索前去重, 不行, 因为三元组中可能有重复

	7. 首先必然有一个循环, 然后需要确定另外两个数, 其中一个数一旦确定最后一个也就确定了
	8. 为了不重复, 另外两个数必然要到最外层没有循环过的范围内找
	9. 没有循环过的范围内也可能找到重复解, 所以在次循环中需要查询上次有没有找到过
	10. 最外循环需要跳过之前已经搜索到的值, 因为之前的搜索中, 搜索其余两个数的范围比后面的都要大, 理论上会全部找全

	给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。



示例 1：

输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/3sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package leetcode

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	res := [][]int{}
	m_i := map[int]int{}
	m_j := map[int]int{}
	m_z := map[int]int{}
	m_p := map[int]map[int]int{}
	val := 0
	for out_index, i := range nums {

		if _, ok := m_i[i]; ok {
			continue
		}
		// fmt.Println("out: ", i)
		m_i[i] = 0
		if m, ok := m_p[i]; ok {
			m_j = m
		} else {
			m_j = map[int]int{}
		}
		m_z = map[int]int{}
		for _, j := range nums[out_index+1:] {

			if _, ok := m_j[j]; ok {
				continue
			}
			val = -i - j
			if _, ok := m_z[val]; ok {
				res = append(res, []int{i, val, j})
				m_j[j], m_j[val] = 0, 0
				if m := m_p[j]; m == nil {
					m_p[j] = map[int]int{val: 0}
				} else {
					m[val] = 0
				}
				if m := m_p[val]; m == nil {
					m_p[val] = map[int]int{j: 0}
				} else {
					m[j] = 0
				}
			} else {
				m_z[j] = 0
			}
		}

	}
	return res
}

func threeSum0(nums []int) [][]int {
	// 与第一种基本没有什么差别
	sort.Ints(nums)
	res := [][]int{}
	m_i := map[int]int{}
	m_j := map[int]int{}
	m_z := map[int]int{}
	val := 0
	for out_index, i := range nums {

		if _, ok := m_i[i]; ok {
			continue
		}
		// fmt.Println("out: ", i)
		m_i[i] = 0
		m_j = map[int]int{}
		m_z = map[int]int{}
		for _, j := range nums[out_index+1:] {

			if _, ok := m_j[j]; ok {
				continue
			}
			val = -i - j
			if _, ok := m_z[val]; ok {
				res = append(res, []int{i, val, j})
				m_j[j], m_j[val] = 0, 0
			} else {
				m_z[j] = 0
			}
		}

	}
	return res
}

// 解法一 最优解，双指针 + 排序
// 这个解法最基本的认识是, 排序后, 三个解一定是按先后顺序出现在结果中
// 所以采用先取中间, 在首尾相夹的办法来解决
// 分阵营(正负零)对比这个而言蠢爆了
func threeSum1(nums []int) [][]int {
	sort.Ints(nums)
	result, start, end, index, addNum, length := make([][]int, 0), 0, 0, 0, 0, len(nums)
	for index = 1; index < length-1; index++ {
		start, end = 0, length-1
		if index > 1 && nums[index] == nums[index-1] {
			start = index - 1
		}
		for start < index && end > index {
			if start > 0 && nums[start] == nums[start-1] {
				start++
				continue
			}
			if end < length-1 && nums[end] == nums[end+1] {
				end--
				continue
			}
			addNum = nums[start] + nums[end] + nums[index]
			if addNum == 0 {
				result = append(result, []int{nums[start], nums[index], nums[end]})
				start++
				end--
			} else if addNum > 0 {
				end--
			} else {
				start++
			}
		}
	}
	return result
}

// 解法二
func threeSum2(nums []int) [][]int {
	res := [][]int{}
	counter := map[int]int{}
	for _, value := range nums {
		counter[value]++
	}

	uniqNums := []int{}
	for key := range counter {
		uniqNums = append(uniqNums, key)
	}
	sort.Ints(uniqNums)

	for i := 0; i < len(uniqNums); i++ {
		if (uniqNums[i]*3 == 0) && counter[uniqNums[i]] >= 3 {
			res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[i]})
		}
		for j := i + 1; j < len(uniqNums); j++ {
			if (uniqNums[i]*2+uniqNums[j] == 0) && counter[uniqNums[i]] > 1 {
				res = append(res, []int{uniqNums[i], uniqNums[i], uniqNums[j]})
			}
			if (uniqNums[j]*2+uniqNums[i] == 0) && counter[uniqNums[j]] > 1 {
				res = append(res, []int{uniqNums[i], uniqNums[j], uniqNums[j]})
			}
			c := 0 - uniqNums[i] - uniqNums[j]
			if c > uniqNums[j] && counter[c] > 0 {
				res = append(res, []int{uniqNums[i], uniqNums[j], c})
			}
		}
	}
	return res
}
