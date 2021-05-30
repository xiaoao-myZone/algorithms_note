/*

要点:
	1. 找出最接近目标值的答案, 假设只存在唯一解
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 2, 1, -4}
	target := 1
	ret := threeSumClosest(nums, target)
	fmt.Println("The result is ", ret)
}

func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	return 0
}
