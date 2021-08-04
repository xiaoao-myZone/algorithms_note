package base

func SearchInsertPoint(nums []int, target int) int {
	// nums is ascent, and if target is largest, the return will be -1
	// target is not in nums, or
	//
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
