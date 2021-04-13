#/usr/bin/env python
# -*- coding: utf-8 -*-
"""
执行用时：48 ms, 在所有 Python 提交中击败了30.95% 的用户
内存消耗：13 MB, 在所有 Python 提交中击败了51.58% 的用户

"""

"""
pre:
    1. 对于`3`与`32`,一定是`32`排在`3`后面
    2. 如果`3`与`34`,一定时`34`排在`3`之前
    3. `3`与`33`随便

post:
    1.

"""
class Solution(object):
    def largestNumber(self, nums):
        """给定一组非负整数 nums，重新排列每个数的顺序（每个数不可拆分）使之组成一个最大的整数
        :type nums: List[int]
        :rtype: str
        """

        nums = [str(i) for i in nums]
        length = len(nums)
        for i in range(length-1):
            for j in range(0, length-i-1):
                a = nums[j]
                b = nums[j+1]
                if a+b<b+a:
                    nums[j], nums[j+1] = b, a
        while nums[0] == '0' and len(nums) is not 1:
            nums.pop(0)
        return "".join(nums)


if __name__ == "__main__":
    from timeit import timeit
    #n = [432,43243]
    n = [3,32,34,9,5]
    #from ipdb import set_trace; set_trace()
    ret = Solution().largestNumber(n)
    print(ret)
    # import timeit
    # timer = timeit.Timer("Solution().largestNumber(%s)" % n, "from __main__ import Solution")
    # print("timeConsume: %s" % timer.timeit(10))