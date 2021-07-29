/*
Score:


Points:
1.

Thoughts:
1. 要么用极值的概念, 要么用栈的概念
2. 后者代码更少一点
3. 这题的巧妙之处在于理解水是如何填满水的槽， 栈的思维可以解决从一侧按顺序填满的过程， 而不是天降暴雨
4. 栈中存储的对象至少是一个二元数组
5. 不一定， 栈中的元素在一定是连续的，只用知道初始索引和长度即可
6. 本质是双(三)指针



给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。



示例 1：

输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。

*/
package leetcode

type block struct {
	height int
	index  int
}

func trap(height []int) int {
	// origin 到 left 之间是一个非升序列
	res, origin, left := 0, 0, 0
	for right := 1; right < len(height); right++ {
		if height[right] > height[right-1] { // 蓄水次上升过程，需要结算
			if height[right] > height[origin] {
				shorter = height[origin]
			} else {
				shorter = height[right]
			}
			for index := right - 1; index >= 0 && height[index] <= shorter; index-- {
				res += shorter - height[index]
				height[index] = shorter
			}
			if height[origin] == shorter {
				origin = right

			}
			// for height[left] < height[right] {
			// 	res += (shorter - height[left]) * distance
			// 	left--
			// 	if left == 0 {
			// 		left = right
			// 		break
			// 	}
			// }
		} else { //下降或者平流过程
			left++
		}

	}
	return res
}

// func trap(height []int) int {
// 	res, tmp, count, distance := 0, make([]int, len(height)), 0, 0
// 	for index, num := range height {
// 		fmt.Println("******: ", num)
// 		//第一个值为0该如何
// 		if num > tmp[count] {

// 			if count == 0 { // 如果count为零表示已经结算完成
// 				tmp[0] = num
// 				fmt.Println("1 ", tmp[:count+1])
// 			} else { //否则是从谷底往上升
// 				// if tmp[0] > num {
// 				// 	shorter = num
// 				// } else {
// 				// 	shorter = tmp[0]
// 				// }
// 				for i := 0; ; i++ { //因为tmp的入栈方式， 使它不可能出现count=-1的情况(不是)
// 					distance++
// 					if tmp[count] >= num {
// 						break
// 					}
// 					fmt.Println("2 ", tmp[:count+1])
// 					res += (num - tmp[count]) * (i + distance)
// 					if count == 0 {
// 						tmp[0] = num
// 						distance = 0
// 						break
// 					}
// 					count--

// 				}

// 			}

// 		} else { //如果小于说明可以入栈， 是一个往谷底下探的过程
// 			if tmp[0] > num {
// 				count++
// 				tmp[count] = num
// 				fmt.Println("3 ", tmp[:count+1])
// 			}

// 		}
// 		fmt.Println("res ", res)
// 	}
// 	fmt.Println("last res ", res)

// 	return res
// }

// func trap(height []int) int {
// 	res, tmp, count, shorter := 0, make([]int, len(height)), 0, 0
// 	fmt.Println("初始: ", tmp[:count+1])
// 	for _, num := range height {
// 		fmt.Println("******: ", num)
// 		//第一个值为0该如何
// 		if num > tmp[count] {

// 			if count == 0 { // 如果count为零表示已经结算完成
// 				tmp[0] = num
// 				fmt.Println("1 ", tmp[:count+1])
// 			} else { //否则是从谷底往上升
// 				if tmp[0] > num {
// 					shorter = num
// 				} else {
// 					shorter = tmp[0]
// 				}
// 				for { //因为tmp的入栈方式， 使它不可能出现count=-1的情况(不是)
// 					if tmp[count] >= num {
// 						break
// 					}
// 					fmt.Println("2 ", tmp[:count+1])
// 					res += shorter - tmp[count]
// 					if count == 0 {
// 						tmp[0] = num
// 						break
// 					}
// 					count--

// 				}

// 			}

// 		} else { //如果小于说明可以入栈， 是一个往谷底下探的过程
// 			if tmp[0] > num {
// 				count++
// 				tmp[count] = num
// 				fmt.Println("3 ", tmp[:count+1])
// 			}

// 		}
// 		fmt.Println("res ", res)
// 	}
// 	fmt.Println("last res ", res)
// 	return res

// }
