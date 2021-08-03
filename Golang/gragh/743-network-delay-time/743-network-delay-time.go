/*
Score:


Points:
1.

Thoughts:
1. 不可以用分层传递法，来确定最短传播时间
2. 应该遍历每一条路， 对其求和
3. 应该用回溯,迭代的思路
4. 将每一个点的上一个点积累的时间加到当前
5. 应该将每一个末端节点的累计时间放入一个map里， 因为末端节点的入度不一定为1
6. 可能两个点相互指向， 这可能就是n的作用了

有 n 个网络节点，标记为 1 到 n。

给你一个列表 times，表示信号经过 有向 边的传递时间。 times[i] = (ui, vi, wi)，其中 ui 是源节点，vi 是目标节点， wi 是一个信号从源节点传递到目标节点的时间。

现在，从某个节点 K 发出一个信号。需要多久才能使所有节点都收到信号？如果不能使所有节点收到信号，返回 -1 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/network-delay-time
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。


*/
package leetcode

import "fmt"

func networkDelayTime(times [][]int, n int, k int) int {
	fmt.Println("1*******")
	tmp := map[int]map[int][]int{}                                                      //map[int]map[int][]int{} //map[int][][]int{}
	res, left, right, nodes, last_left, res_tmp := 0, 0, 0, [][]int{}, 0, map[int]int{} //make([][]int, n)
	for _, data := range times {
		if _, ok := tmp[data[0]]; ok {
			tmp[data[0]][data[1]] = data
			//tmp[data[0]] = append(tmp[data[0]], data)
		} else {
			tmp[data[0]] = map[int][]int{data[1]: data}
			//tmp[data[0]] = [][]int{data}
		}
	}
	fmt.Println("2*******")
	for _, v := range tmp[k] {
		nodes = append(nodes, v)
		right++
	}
	if right == 0 {
		fmt.Println("5*******: -1")
		return -1
	}
	fmt.Println("3*******init_node: ", nodes)
	//nodes[0] = times[k]
	for right < len(times) {

		if left == right {
			fmt.Println("5*******: -1： ", left, right)
			return -1
		}
		last_left = left // 存上一次的left值，为最后核算做准备
		fmt.Println("3.1*******: ", len(nodes), left, right)
		//寻找下一层传播点集
		for _, num := range nodes[left:right] {
			left++
			fmt.Println("3.2******* 选中: ", num)
			if datas, ok := tmp[num[1]]; ok {
				fmt.Println("3.3******* 下探: ", datas)
				// k 为num[1]的下游节点
				for _, v := range datas {
					nodes = append(nodes, v)
					right++
					// 找到节点k的上一个节点nums[1]/v[0]，并将它的累积时间加到自身
					// 对于第二层节点， 第一层是找不到的
					fmt.Println("$$$", v, num[2])
					v[2] += num[2]
					fmt.Println("3.4*******: v update", v, right)
					fmt.Println("3.4*******")
				}

			} else {
				//未到尽头，但是时间不见得短, 保留
				// 取最短时间
				if v, ok := res_tmp[num[1]]; ok {
					if v > num[2] {
						res_tmp[num[1]] = num[2]
					}
				} else {
					res_tmp[num[1]] = num[2]
				}

			}
			// TODO python里面的for else该如何实现？
		}
	}
	fmt.Println("4*******： ", res_tmp)

	for _, num := range nodes[last_left:] {
		fmt.Println("the last: ", num)
		// 取最短时间
		if v, ok := res_tmp[num[1]]; ok {
			if v > num[2] {
				res_tmp[num[1]] = num[2]
			}
		} else {
			res_tmp[num[1]] = num[2]
		}
	}
	for _, v := range res_tmp {
		if res < v {
			res = v
		}
	}
	fmt.Println("5*******: ", res)
	return res

}
