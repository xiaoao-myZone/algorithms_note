/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00% 的用户
内存消耗：2.3 MB, 在所有 Go 提交中击败了88.24% 的用户
笑死
有效数字: 整数, 小数, 科学计数法

["2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3", "3e+7", "+6e-1", "53.5e93", "-123.456e789"]
["abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53"]

*/
package main

import "fmt"

func main() {
	right := []string{"2", "0089", "-0.1", "+3.14", "4.", "-.9", "2e10", "-90E3", "3e+7", "+6e-1", "53.5e93", "-123.456e789"}
	wrong := []string{"abc", "1a", "1e", "e3", "99e2.5", "--6", "-+3", "95a54e53", ".", ".-4"}
	for _, i := range right {
		ret := isNumber(i)
		fmt.Printf("%t\t", ret)
	}
	fmt.Printf("\n")
	for _, i := range wrong {
		ret := isNumber(i)
		fmt.Printf("%t\t", ret)
	}
}

func isNumber(s string) bool {
	// 有过符号, 有过科学计数法, 有过小数点, 有过数字
	signed, isSci, isDecimal, digit := false, false, false, false
	for _, i := range s {
		if '0' <= i && i <= '9' {
			if !digit { //记得遇到过一个数字
				digit = true
			}
		} else if i == '+' || i == '-' {
			if signed || digit || isDecimal { //之前不可以有符号 不可以有小数点 不可以有数字
				return false
			} else {
				signed = true
			}
		} else if i == 'e' || i == 'E' {
			if isSci || !digit { // 之前不可以有e/E, 不可以没有数字
				return false
			} else { //之后的数字需要重新判断是否是整数
				isSci, signed, isDecimal, digit = true, false, false, false
			}
		} else if i == '.' {
			if isDecimal { //之前不可以有小数点
				return false
			} else {
				isDecimal = true
			}
		} else {
			return false
		}
	}

	if isSci {
		if !digit || isDecimal { //科学计数标志e后面需要时整数
			return false
		}
	}
	if isDecimal {
		if !digit { //小数点后面至少有一个数字
			return false
		}
	}
	return true
}
