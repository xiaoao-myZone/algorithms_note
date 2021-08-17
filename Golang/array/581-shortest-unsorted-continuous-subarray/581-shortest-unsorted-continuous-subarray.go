/*
Score:
执行用时：132 ms, 在所有 Go 提交中击败了5.93% 的用户
内存消耗：6.5 MB, 在所有 Go 提交中击败了48.62% 的用户
qvq
执行用时：24 ms, 在所有 Go 提交中击败了92.14% 的用户
内存消耗：6.4 MB, 在所有 Go 提交中击败了79.56% 的用户
Points:
1.

Thoughts:
1. 如果第一个数不是最小值， 那么这个序列一定是从头开始
2. 如果最后一个数不是最大值， 那么这个序列一定在尾部结束
3. 升序的结果是唯一的
4. 反过头来想， 可以寻找保序的字段，再反推无序的部分
5. 这样需要二分法查找插入位置
6. 存在相同的数值
7. 排序都只要O(nlogn)，第一种解法居然要O(n**2)显然不正常

给你一个整数数组 nums ，你需要找出一个 连续子数组 ，如果对这个子数组进行升序排序，那么整个数组都会变为升序排序。

请你找出符合题意的 最短 子数组，并输出它的长度。

示例 1：

输入：nums = [2,6,4,8,10,9,15]
输出：5
解释：你只需要对 [6, 4, 8, 10, 9] 进行升序排序，那么整个表都会变为升序排序。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/shortest-unsorted-continuous-subarray
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
package leetcode

// import "./Golang/base"

func findUnsortedSubarray(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	left, right := 0, 0
	for i := 0; i < len(nums); i++ {
		for _, tmp := range nums[i+1:] {
			if nums[i] > tmp {
				left = i
				i = len(nums) //break outer loop
				break
			}
		}
	}

	for i := len(nums) - 1; i > left; i-- {
		for _, tmp := range nums[left:i] {
			if nums[i] < tmp {
				right = i
				i = 0 //break outer loop
				break
			}
		}
	}
	if left == right {
		return 0
	} else {
		return right - left + 1
	}

}

func findUnsortedSubarray1(nums []int) int {
	left, right, maxium, minium := 0, len(nums)-1, 0, 0
	for ; left < len(nums)-1; left++ {
		// 左边界至少在这里
		if nums[left] > nums[left+1] {
			maxium = nums[left]
			break
		}

	}
	for ; right > left; right-- {
		// 右边界至少在这里
		if nums[right-1] > nums[right] {
			minium = nums[right]
			break
		}

	}
	if left == right {
		return 0
	}
	//找到这个区间的最小值与最大值
	for i := left + 1; i <= right-1; i++ {
		if nums[i] > maxium {
			maxium = nums[i]
		} else {
			if nums[i] < minium {
				minium = nums[i]
			}
		}
	}
	for i := 0; i < left; i++ {
		if nums[i] > minium {
			left = i
			break
		}
	}
	for i := len(nums) - 1; i > right; i-- {
		if nums[i] < maxium {
			right = i
			break
		}
	}

	return right - left + 1
}

func SearchInsertPoint(nums []int, target int) int {
	// nums is ascent, and if target is largest, the return will be -1
	// target is not in nums, or
	if len(nums) == 0 {
		return 0
	}
	res := len(nums) / 2 //偏右
	tmp := 0
	if nums[0] > target {
		return 0
	}
	if nums[len(nums)-1] < target {
		return len(nums)
	}
	//res != 0 && res != len(nums)
	for {
		if nums[res] > target {
			if nums[tmp] < target && res-tmp == 1 {
				return res
			} else {
				tmp = res
				res /= 2
			}

		} else if nums[res] < target {
			if nums[tmp] > target && tmp-res == 1 {
				return tmp
			} else {
				tmp = res
				res = (len(nums) + res) / 2
			}

		} else {
			return -1 //target in nums
		}
	}
	// return -1 //target in nums
}
