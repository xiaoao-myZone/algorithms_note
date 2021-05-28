/*


要求:返回一个数组中三个不同位置且其和等于0的所有解
分析:
	1. 貌似最多只能做到O(n2)
	2. 可以两两组合和, 用0减去, 在列表中搜索该值
*/
package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}
	ret := threeSum(nums)
	fmt.Println("The result is ", ret)
}

func threeSum(nums []int) [][]int {
	return [][]int{}
}
