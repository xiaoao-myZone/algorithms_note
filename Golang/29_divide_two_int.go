/*
Score:
执行用时：0 ms, 在所有 Go 提交中击败了100.00% 的用户
内存消耗：2.4 MB, 在所有 Go 提交中击败了12.89% 的用户

Points:
1. 不许用编程语言自带的乘法, 除法以及求余运算
2. 向0取整
3. 范围在正负2^32以内

Thoughts:
1. 减法  循环
2. 兼容负数的情况
3. ####可以借鉴对人类来说是笨办法的方法来求解, 比如选一个结果反推除数
4. 充分利用位运算
5. 从最大数开始, 因为从1开始, 位移不方便补全
6. 将结果看做[a0, a1, a2, a3, ..., an][1, 2, 4, ..., 2^n]T
7. 找到2^n后, 应当将dividend- divisor*2^n作为后续的迭代求解
8. 而divisor*2^n就是divisor<<n
9. 这种办法和小学除法的竖式运算没有区别, 把进制换成任意数, 也都是一样的
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	dividend := 7 //-2147483648
	divisor := 8

	ret := divide1(dividend, divisor)
	fmt.Println("The result is ", ret)
}

func divide(dividend int, divisor int) int {
	count := 0
	is_neg := true
	if dividend >= 0 && divisor > 0 || dividend <= 0 && divisor < 0 {
		is_neg = false
	}
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}
	for dividend >= divisor {
		dividend -= divisor
		if is_neg {
			count--
		} else {
			count++
		}

	}
	// if is_neg {
	// 	count = -count
	// }
	if count > math.MaxInt32 {
		count = math.MaxInt32
	}
	fmt.Println(math.MaxInt32) // 2147483647
	fmt.Println(math.MinInt32) // -2147483648
	return count

}

func divide1(dividend int, divisor int) int {
	bits, res, isNegtive := 31, 0, false
	if dividend >= 0 && divisor >= 0 || dividend <= 0 && divisor <= 0 {
		isNegtive = true
	}
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}
	fmt.Println(dividend, divisor, isNegtive)
	for {
		if dividend < divisor {
			break
		}
		for {
			if dividend>>bits >= divisor {
				break
			}
			bits--
			// fmt.Println(bits)

		}
		res += 1 << bits
		// fmt.Println(res)
		dividend = dividend - divisor<<bits

	}
	if isNegtive {
		if res > math.MaxInt32 {
			return math.MaxInt32
		} else {
			return res
		}

	} else {
		return -res
	}

}
