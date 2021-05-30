/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00% 的用户
内存消耗：2.1 MB, 在所有 Go 提交中击败了99.83% 的用户

>>1 和 /2感觉差不多
*/
package main

import "fmt"

func main() {
	num := 8
	ret := isPowerOfTwo(num)
	fmt.Println("The result is ", ret)
}

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	for n%2 == 0 {
		n = n >> 1
		fmt.Println(n)
	}

	if n == 1 {
		return true
	} else {
		return false
	}
}
