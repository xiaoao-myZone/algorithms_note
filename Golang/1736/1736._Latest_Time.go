/*
Score:
执行用时：0 ms, 在所有 Go 提交中击败了100.00% 的用户
内存消耗：1.9 MB, 在所有 Go 提交中击败了76.39% 的用户

Points:
1.

Thoughts:
1.



给你一个字符串 time ，格式为 hh:mm（小时：分钟），其中某几位数字被隐藏（用 ? 表示）。

有效的时间为 00:00 到 23:59 之间的所有时间，包括 00:00 和 23:59 。

替换 time 中隐藏的数字，返回你可以得到的最晚有效时间。

*/

package leetcode

func maximumTime(time string) string {
	for index, val := range time {
		if val == '?' {
			switch index {
			case 0:
				if time[1] < '4' || time[1] == '?' {
					time = "2" + time[1:]
				} else {
					time = "1" + time[1:]
				}
			case 1:
				if time[0] < '2' {
					time = time[:1] + "9" + time[2:]
				} else {
					time = time[:1] + "3" + time[2:]
				}

			case 3:
				time = time[:3] + "5" + time[4:]
			case 4:
				time = time[:4] + "9"

			}

		}
	}
	return time
}
