/*

除了解的长度, 规则与第15题相同
是否可以当做进行n-3个三数之和
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0
	ret := fourSum(nums, target)
	fmt.Println("The result is ", ret)
}

// TODO
func fourSum(nums []int, target int) [][]int {
	ret := [][]int{}
	sort.Ints(nums)
	return ret
}
