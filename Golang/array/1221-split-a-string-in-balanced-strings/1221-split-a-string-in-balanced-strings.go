/*
Score:
执行用时：0 ms, 在所有 Go 提交中击败了100.00% 的用户
内存消耗：1.9 MB, 在所有 Go 提交中击败了59.74% 的用户

Points:
1.

Thoughts:
1. 不可以用指针的方法， 因为中间穿插问题无法解决


在一个 平衡字符串 中，'L' 和 'R' 字符的数量是相同的。

给你一个平衡字符串 s，请你将它分割成尽可能多的平衡字符串。

注意：分割得到的每个字符串都必须是平衡字符串。

返回可以通过分割得到的平衡字符串的 最大数量 。



示例 1：

输入：s = "RLRRLLRLRL"
输出：4
解释：s 可以分割为 "RL"、"RRLL"、"RL"、"RL" ，每个子字符串中都包含相同数量的 'L' 和 'R' 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/split-a-string-in-balanced-strings
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


*/
package leetcode

// import "fmt"

func balancedStringSplit(s string) int {
	ret, Rcount, Lcount := 0, 0, 0
	for _, word := range s {
		if word == 'R' {
			Rcount++
		} else {
			Lcount++
		}
		if Rcount == Lcount {
			ret += 1
		}
	}
	return ret
}
