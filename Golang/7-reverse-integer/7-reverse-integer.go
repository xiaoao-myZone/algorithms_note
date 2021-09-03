/*
Score:
执行用时：0 ms, 在所有 Go 提交中击败了100.00% 的用户
内存消耗：2.1 MB, 在所有 Go 提交中击败了99.97% 的用户
通过测试用例：1032 / 1032

Points:
1.

Thoughts:
1. 注意溢出
2. 可否用>>来加速（貌似不方便， 做除法还可以）


给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。

如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。
假设环境不允许存储 64 位整数（有符号或无符号）。



示例 1：

输入：x = 123
输出：321

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-integer
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
package leetcode

import (
	"math"
)

func reverse(x int) int {
	ret := 0
	if x < 0 {
		ret = -_reverse(-x)
		if ret < math.MinInt32 {
			return 0
		} else {
			return ret
		}
	} else {
		ret = _reverse(x)
		if ret > math.MaxInt32 {
			return 0
		} else {
			return ret
		}
	}
}

func _reverse(x int) int {
	ret, remain := 0, 0
	for x > 9 {
		remain = x % 10
		x /= 10
		ret += remain
		ret *= 10
	}
	ret += x
	return ret
}
