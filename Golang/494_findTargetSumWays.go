/*

通过在数列数字前添加"+"或"-"使其生成的表达式的值等于目标值, 返回这样的解的数目
有点像24点游戏

思考:
	1. 应该先排序, 然后搜索(没有必要)
	2. 须知, 如果再创建一个slice, 存放1和-1, 经过排列组合一共有2^n种情况
	3. 但是2^n种情况,并不能预测其大小, 比如[1,1,4,5]解[-1,-1,1,1]不比[1,1,-1,1]小
	4. 另外, 这些组合中非只有唯一解, 需要迭代
	5. 如果当前值比target小, 那么继续增加-1的数量是不可取的
	6. 暂时没想出好法子, 这题考得应该是迭代
*/
package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 1}
	target := 0
	ret := findTargetSumWays(nums, target)
	fmt.Println("The result is ", ret)
}

func findTargetSumWays(nums []int, target int) int {
	//默认初始解是[1]*n
	ret := 0
	if len(nums) == 0 {
		if target == 0 {
			return 1
		} else {
			return 0
		}
	}
	the_sum := sum(nums)
	fmt.Println(the_sum, target, nums)
	if the_sum < target {
		return ret
	} else if the_sum > target {
		for i := range nums {
			i += 1 //Nothing, just to make compile successfully
		}

	} else {
		ret++ // 这样势必包含很多重复的判断
		for _, i := range nums {
			if i == 0 {
				ret *= 2
			}
		}
		return ret
	}
	return ret
}

func sum(nums []int) int {
	ret := 0
	for _, i := range nums {
		ret += i
	}
	return ret
}
