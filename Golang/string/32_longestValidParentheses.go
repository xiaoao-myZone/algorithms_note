/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00% 的用户
内存消耗：3.4 MB, 在所有 Go 提交中击败了40.07% 的用户

分析:
1. 每遇到一个')'就开始局部结算, 结果是0或1
2. 连续的1可以累积
3. 一个区域累积完成后, 需要与上一个的相比较
4. 只有当(有盈余的时候才有必要接着做
5. 另一个关键在于, 积累的(能不能消耗完, 需要到结尾才知道, 不可避免的需要用到栈
*/
package main

import "fmt"

func main() {
	test := []*string{}
	a := "()(())"
	test = append(test, &a)
	b := "()(()"
	test = append(test, &b)
	c := "))))((()(("
	test = append(test, &c)
	fmt.Println("test: ", test)
	for _, n := range test {
		fmt.Println("s: ", *n)
		ret := longestValidParentheses(*n)
		fmt.Println("The result is ", ret)
	}

}

func longestValidParentheses(s string) int {
	ret := 0
	nums := []int{0}
	for _, i := range s {
		if i == '(' {
			nums = append(nums, 0) //吃进(, 记录在哪里吃进
			fmt.Println("吃进: ", nums)
		} else {
			if len(nums) > 1 {
				//吐出(
				nums[len(nums)-2] = 2 + nums[len(nums)-2] + nums[len(nums)-1]
				nums = nums[:len(nums)-1] // pop
				fmt.Println("缩卷: ", nums)
			} else { //此时nums的长度必然是1, 因为)盈余
				fmt.Println("结算: ", nums)
				if nums[0] > ret {
					ret = nums[0]
				}
				nums[0] = 0
			}
		}
	}
	// 最后长度只会是1或者2 (不对)
	fmt.Println(nums)
	for _, i := range nums {
		if i > ret {
			ret = i
		}
	}

	return ret
}
