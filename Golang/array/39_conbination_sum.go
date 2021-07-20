/*
Score:


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
1.
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

}
