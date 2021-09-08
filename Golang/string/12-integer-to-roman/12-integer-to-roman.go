/*
Score:
执行用时：4 ms, 在所有 Go 提交中击败了94.83% 的用户
内存消耗：3.5 MB, 在所有 Go 提交中击败了24.34% 的用户

Points:
1. 1 <= num <= 3999

Thoughts:
1.

罗马数字包含以下七种字符： I， V， X， L，C，D 和 M。

字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000

例如， 罗马数字 2 写做 II ，即为两个并列的 1。12 写做 XII ，即为 X + II 。 27 写做  XXVII, 即为 XX + V + II 。

通常情况下，罗马数字中小的数字在大的数字的右边。但也存在特例，例如 4 不写做 IIII，而是 IV。数字 1 在数字 5 的左边，所表示的数等于大数 5 减小数 1 得到的数值 4 。同样地，数字 9 表示为 IX。这个特殊的规则只适用于以下六种情况：

    I 可以放在 V (5) 和 X (10) 的左边，来表示 4 和 9。
    X 可以放在 L (50) 和 C (100) 的左边，来表示 40 和 90。
    C 可以放在 D (500) 和 M (1000) 的左边，来表示 400 和 900。

给你一个整数，将其转为罗马数字。



示例 1:

输入: num = 3
输出: "III"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/integer-to-roman
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


*/
package leetcode

// import "fmt"

func intToRoman(num int) string {
	ret := ""
	for bit := 1; num > 0; bit, num = bit+1, num/10 {
		ret = forOneBit(num%10, bit) + ret
	}
	return ret
}

func forOneBit(num int, bit int) string {
	ret, combines := "", []string{}
	if bit == 1 {
		combines = []string{"I", "V", "X"}
	} else if bit == 2 {
		combines = []string{"X", "L", "C"}
	} else if bit == 3 {
		combines = []string{"C", "D", "M"}
	} else {
		combines = []string{"M"}
	}
	if num == 9 {
		return combines[0] + combines[2]
	}
	if num == 4 {
		return combines[0] + combines[1]
	}
	if num > 4 {
		ret = combines[1]
		num -= 5
	}
	for i := 0; i < num; i++ {
		ret += combines[0]
	}
	return ret
}
