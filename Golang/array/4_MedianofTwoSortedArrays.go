/*
执行用时：12 ms, 在所有 Go 提交中击败了94.06% 的用户
内存消耗：6.5 MB, 在所有 Go 提交中击败了5.35% 的用户
*/

/*
1. 需不需要先合并?
2. 最好能一遍合并一边找
3. 中位数的位数可以提前找
4. 合并是创建一个新的分片还是续用原来的(没必要专门用一个分片来做这个)
5. 奇偶统一很难
*/
package main

import "fmt"

func main() {
	num1 := []int{}
	num2 := []int{1}
	// var m float64
	// m = (2 + 5) / 2.0
	// fmt.Println(num1, num2, m)
	ret := findMedianSortedArrays(num1, num2)
	fmt.Printf("%f\n", ret)

}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	i := len(nums1) - 1
	j := len(nums2) - 1
	// fmt.Println(i, j)
	length := i + j + 1
	middle := length / 2 // 注意奇偶
	sum := []int{}
	for i >= 0 && j >= 0 {
		m := nums1[i]
		n := nums2[j]
		if m > n {
			sum = append(sum, m)
			i -= 1
		} else {
			sum = append(sum, n)
			j -= 1
		}
	}
	// fmt.Println(i, j)
	for i >= 0 {
		sum = append(sum, nums1[i])
		i -= 1
	}
	for j >= 0 {
		sum = append(sum, nums2[j])
		j -= 1
	}
	// sum = append(sum, nums1[:i+1]...)
	// sum = append(sum, nums2[:j+1]...)
	// fmt.Println(sum, middle, length, sum[middle], sum[length-middle])
	return float64((sum[middle] + sum[length-middle])) / 2
}

func findMedianSortedArrays_1(nums1 []int, nums2 []int) float64 {
	//真的妙
	// 假设 nums1 的长度小
	/*
		1. 未必一定要在一个融合两个数组的数组中才能找到中线
		2. 中线只是一个阵营的划分, 中线两边必然都由两个原数组组成
		3. 根据奇偶, 找到中线两边的最大值与最小值
	*/
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}
	low, high, k, nums1Mid, nums2Mid := 0, len(nums1), (len(nums1)+len(nums2)+1)>>1, 0, 0
	for low <= high {
		// nums1:  ……………… nums1[nums1Mid-1] | nums1[nums1Mid] ……………………
		// nums2:  ……………… nums2[nums2Mid-1] | nums2[nums2Mid] ……………………
		nums1Mid = low + (high-low)>>1                           // 本质是(low + high)>>1                 // 分界限右侧是 mid，分界线左侧是 mid - 1
		nums2Mid = k - nums1Mid                                  // 也是
		if nums1Mid > 0 && nums1[nums1Mid-1] > nums2[nums2Mid] { // nums1 中的分界线划多了，要向左边移动
			high = nums1Mid - 1
		} else if nums1Mid != len(nums1) && nums1[nums1Mid] < nums2[nums2Mid-1] { // nums1 中的分界线划少了，要向右边移动
			low = nums1Mid + 1 //这里重复计算nums1的长度没有问题吗?难道是go内部优化了? TODO
		} else {
			// 找到合适的划分了，需要输出最终结果了
			// 分为奇数偶数 2 种情况
			break
		}
	}
	midLeft, midRight := 0, 0
	if nums1Mid == 0 {
		midLeft = nums2[nums2Mid-1]
	} else if nums2Mid == 0 {
		midLeft = nums1[nums1Mid-1]
	} else {
		midLeft = max(nums1[nums1Mid-1], nums2[nums2Mid-1])
	}
	if (len(nums1)+len(nums2))&1 == 1 {
		return float64(midLeft)
	}
	if nums1Mid == len(nums1) {
		midRight = nums2[nums2Mid]
	} else if nums2Mid == len(nums2) {
		midRight = nums1[nums1Mid]
	} else {
		midRight = min(nums1[nums1Mid], nums2[nums2Mid])
	}
	return float64(midLeft+midRight) / 2
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
