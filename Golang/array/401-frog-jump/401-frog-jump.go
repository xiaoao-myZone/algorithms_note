/*
Score:
timeout fail

Points:
1.

Thoughts:
1. 直观的想法是， 不知道每一步的走法会对后面造成什么影响， 直到走投无路
2. 然而， 这里每步最多有三个选择， 所以时间复杂度是O(3**n)
3. 但是数列的长度限制了这种复杂度， 比如，可能跳到第n个石头有多种方法， 但是从n到下一个情况是较少的
4. 可以尝试反其道而行之，从尾部开始，先假设步数， 在判断之前的动作(不行)
5. 走n步最多可以走到n(n+1)/2， 数列长度为l， n=l， 带入， 即为队尾数值的最大值l(l-1)/2， 最小值是l-1


一只青蛙想要过河。 假定河流被等分为若干个单元格，并且在每一个单元格内都有可能放有一块石子（也有可能没有）。 青蛙可以跳上石子，但是不可以跳入水中。

给你石子的位置列表 stones（用单元格序号 升序 表示）， 请判定青蛙能否成功过河（即能否在最后一步跳至最后一块石子上）。

开始时， 青蛙默认已站在第一块石子上，并可以假定它第一步只能跳跃一个单位（即只能从单元格 1 跳至单元格 2 ）。

如果青蛙上一步跳跃了 k 个单位，那么它接下来的跳跃距离只能选择为 k - 1、k 或 k + 1 个单位。 另请注意，青蛙只能向前方（终点的方向）跳跃。



示例 1：

输入：stones = [0,1,3,5,6,8,12,17]
输出：true
解释：青蛙可以成功过河，按照如下方案跳跃：跳 1 个单位到第 2 块石子, 然后跳 2 个单位到第 3 块石子, 接着 跳 2 个单位到第 4 块石子, 然后跳 3 个单位到第 6 块石子, 跳 4 个单位到第 7 块石子, 最后，跳 5 个单位到第 8 个石子（即最后一块石子）。

示例 2：

输入：stones = [0,1,2,3,4,8,9,11]
输出：false
解释：这是因为第 5 和第 6 个石子之间的间距太大，没有可选的方案供青蛙跳跃过去。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/frog-jump
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
package leetcode

// import "fmt"

func canCross(stones []int) bool {
	//fmt.Println("**********************")
	sMap, last_index := map[int]int{}, len(stones)-1
	for index, num := range stones {
		sMap[num] = index
	}
	return jump(stones, sMap, 0, 0, last_index)
}

func jump(stones []int, sMap map[int]int, lastStep int, first_index int, last_index int) bool {

	for step, i := lastStep-1, 0; i < 3; i, step = i+1, step+1 {
		if step > 0 {
			next := stones[first_index] + step
			//fmt.Printf("choose: %v start: %v:%v-->%v\n", i, first_index, stones[first_index], next)
			if next_index, ok := sMap[next]; ok {
				if next_index != last_index {
					if res := jump(stones, sMap, step, next_index, last_index); res {
						return true
					} else {
						continue
					}
				} else {
					return true
				}
			} else {
				continue
			}

		} else {
			continue
		}
	}
	return false
}
