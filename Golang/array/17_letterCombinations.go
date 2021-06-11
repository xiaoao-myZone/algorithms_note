/*
执行用时：0 ms, 在所有 Go 提交中击败了100.00% 的用户
内存消耗：2.1 MB, 在所有 Go 提交中击败了14.87% 的用户

TODO 参考halfrost, 降低内存消耗
说明:
	九键键盘一串数字(2-9)可能的字母组合
	将每一次遇到的数字的所有可能与之前的排列组合
Conclusion:
	1. 在迭代计算中, 最大内存占用量是 迭代深度*每调用一次该函数所消耗的内存
*/
package main

import "fmt"

func main() {
	str := "23"
	ret := letterCombinations(str)

	fmt.Println("The result is ", ret)
}

func letterCombinations(digits string) []string {
	ret, tmp := []string{}, []string{}
	//fmt.Println(digits)
	if len(digits) == 0 {
		return ret
	} else {

		ret = do(ret, digits[0]) //没有必要, 因为ret必为空
		tmp = letterCombinations(digits[1:])
		ret = combine(ret, tmp)
		return ret

	}

	// return ret
}
func combine(a []string, b []string) []string {
	if len(a) == 0 || len(b) == 0 {
		return append(a, b...)
	}
	len_a := len(a)
	for _, j := range b {
		for i := 0; i < len_a; i++ {
			a = append(a, a[i]+j)
		}
	}
	return a[len_a:]
}

func do(a []string, s byte) []string {
	var tmp []string
	switch s {
	case '2':
		tmp = []string{"a", "b", "c"}
	case '3':
		tmp = []string{"d", "e", "f"}
	case '4':
		tmp = []string{"g", "h", "i"}
	case '5':
		tmp = []string{"j", "k", "l"}
	case '6':
		tmp = []string{"m", "n", "o"}
	case '7':
		tmp = []string{"p", "q", "r", "s"}
	case '8':
		tmp = []string{"t", "u", "v"}
	case '9':
		tmp = []string{"w", "x", "y", "z"}
	}
	return combine(a, tmp)
}
