/*
Score:


Points:
1. 先找对称轴
2. 已经比对过的结果貌似对后续的比对没有任何利用价值， 所以要重头来， 时间复杂对为O(n^2)不可避免
3. 但是可以记录其长度， 但是后续的最长搜索都没有这个长时， 也就没有了搜索价值

Thoughts:
1. 从中间还是从两边？ 从两边尝试， 即便成功了， 也不能断言之后没有更长的
2. 但是如果从中间尝试使结果足够长， 可以不必继续尝试边缘
3. 如何体现从中间到两边？
4. 如何体现以字符本省为对称轴？又以字符间的间隙为对称轴? 如果倍增长度， 然后通过奇偶性来区分，感觉会
5. 不用遍历到最边缘, 所以最多一共有2*n-3次尝试



*/
package leetcode

// import "fmt"

func longestPalindrome(s string) string {
	ret, tmp, left, right, i, j, i_log, j_log, log := "", "", 0, len(s)/2, len(s)/2, 0, true, false, true
	for i > 0 || j < len(s) {
		// 左边
		if log {
			if i_log {
				left, right = i, i
				compare(left, right, &s)
				i_log = false
			} else {
				i_log = true
				i--
			}
		}
	}
	return ret
}

func compare(i int, j int, s *string) string {
	for i > 0 && j < len(*s) {
		if (*s)[i] != (*s)[j] {
			break
		}
	}
	return
}
