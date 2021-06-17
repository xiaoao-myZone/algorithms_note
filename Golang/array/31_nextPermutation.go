/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00% 的用户
内存消耗：2.5 MB, 在所有 Go 提交中击败了23.08% 的用户

只申请了两个变量, 内存就如此的低?

注意理解"更大"
问题转变成,如何交换两个数字,使得新数字刚好比原来的数字大
1. 如果最后两位出现顺序, 那么交换一下一定时就是答案[1,0,1]
2. 如果不是最后两位, 需要比较前一位 [1, 4, 3, 2] 2, 3 , 4, 1
4. 总之选择将最低两位交换, 并且在此基础上将被交换的最高位后面的位数排序为最小
*/
package main

import (
	"fmt"
	"sort"
)

func main() {

	nums := []int{3, 2, 1} //{2,3,1} {1, 3, 2}
	nextPermutation(nums)

	fmt.Println("The result is ", nums)
}
func nextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}
	second := len(nums) - 1 //后面待交换的最小的数
	first := second - 1
	for first >= 0 {
		if nums[first] < nums[first+1] { // 能进行到这一步, 说明first+1后面都是逆序
			for true {
				if nums[first] < nums[second] {
					nums[first], nums[second] = nums[second], nums[first]
					break
				} else {
					second--
				}
			}
			break
		} else {
			first--
		}

	}
	sort.Ints(nums[first+1:])
}
