/*
Score:

TODO
Points:
给定一个无重复元素的数组 candidates 和一个目标数 target ，找出 candidates 中所有可以使数字和为 target 的组合。

candidates 中的数字可以无限制重复被选取
1. 假定是排序过得数组吧
2. 先搜索一遍目标值所在位置, 逆序搜
3. 如果不在队尾
4. 可能需要化整为零, 使用迭代的方法
5. ####可以借鉴对人类来说是笨办法的方法来求解, 比如选一个结果反推除数
6. 充分利用位运算

结果是 [q1, q2, ..., qn][a1, a2, ..., an]T = t


Thoughts:
1. 粗算时间复杂度是π(t/qi)
2. 优化的要诀是对给定数组排序,然后在特定情况下, 不必进行接下来的搜索


给定一个无重复元素的正整数数组 candidates 和一个正整数 target ，找出 candidates 中所有可以使数字和为目标数 target 的唯一组合。

candidates 中的数字可以无限制重复被选取。如果至少一个所选数字数量不同，则两种组合是唯一的。

对于给定的输入，保证和为 target 的唯一组合数少于 150 个。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combination-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
package main

import (
	"fmt"
)

func main() {
	candidates := []int{2, 3, 6, 7}
	target := 7
	ret := combinationSum(candidates, target)
	fmt.Println("The result is ", ret)
}

func combinationSum(candidates []int, target int) [][]int {
	return [][]int{}
}
