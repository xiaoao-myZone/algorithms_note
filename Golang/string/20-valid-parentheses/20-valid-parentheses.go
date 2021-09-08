/*
Score:
执行用时：0 ms, 在所有 Go 提交中击败了100.00% 的用户
内存消耗：2 MB, 在所有 Go 提交中击败了34.68% 的用户

Points:
1.

Thoughts:
1. 依然不可以用指针， 由于括号的穿插性，比如 (([]){})， 指针只适合对称性非常好的场合
2. 要指针也可以， 需要把已经消掉的括号删掉， 但是字符不允许这样的操作
3. 这种情况下， 可能需要三指针， 或者说有限个指针就可以解决这个问题
4. 这基于相互抵消的括号对是连续的， 那么前面积累的左括号最多只会被分割为两段
5. 试了一下， 需要四个指针， 并且太麻烦了， 而且判断非常多， 未必会快很多


给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

    左括号必须用相同类型的右括号闭合。
    左括号必须以正确的顺序闭合。



示例 1：

输入：s = "()"
输出：true

示例 2：

输入：s = "()[]{}"
输出：true


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
package leetcode

// import "fmt"

func isValid(s string) bool {
	tmp := []rune{}
	for _, val := range s {
		if val == '(' || val == '[' || val == '{' {
			tmp = append(tmp, val)
		} else {
			if len(tmp) == 0 {
				return false
			} else {
				if res := val - tmp[len(tmp)-1]; 0 < res && res < 3 {
					tmp = tmp[:len(tmp)-1]
				} else {
					return false
				}
			}
		}
	}
	if len(tmp) == 0 {
		return true
	} else {
		return false
	}
}
