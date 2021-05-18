#/usr/bin/env python
# -*- coding: utf-8 -*-
"""

"""

"""
[1,5,3,2,10]
pre:
    1. 结果未必会是一个等差为1的数列, 如[1,5,3,2,10], 结果是5+10最大
    2. 如果某一个数特别大, 可能最优解是通过跳步来实现的
    3. 虽然结果不是一个等差数列, 但是, 任何两个数相隔不会超过2 (关键)
    4. 虽然给的是一个数组,但是是环形的, 所以从任意一个元素开始搜索最优值, 计算流程都一样
    5. 根据推断-3, 任选一个,很可能是最优解所跳过的, 所以需要分别尝试一个连在一起的三元组


post:
    1. 使用迭代穷举法时间与空间复杂度指数级qvq

"""
class Solution(object):
    def __init__(self):
        self.ret = 0
    def rob(self, nums):
        """房子首位相连, 每间房(元素)存放一定的现金(元素的值), 不可以偷相邻的房间的钱, 问一次最多偷多少
        :type nums: List[int]
        :rtype: int
        """
        self.len = len(nums)
        length = len(nums)
        if length<4:
            return max(nums)
        self.search(nums[:-1], tmp=0)
        self.search(nums[1:], tmp=0)
        self.search(nums[2:], tmp=0)

        return self.ret
    
    def search(self, nums, tmp):
        if nums:
            tmp+=nums[0]
            self.search(nums[2:], tmp)
            self.search(nums[3:], tmp)
        else:
            if tmp>self.ret:
                self.ret = tmp

if __name__ == "__main__":
    from timeit import timeit
    nums = [226,174,214,16,218,48,153,131,128,17,157,142,88,43,37,157,43,221]#,191,68,206,23,225,82,54,118,111,46,80,49,245,63,25,194,72,80,143,55,209,18,55,122,65,66,177,101,63,201,172,130,103,225,142,46,86,185,62,138,212,192,125,77,223,188,99,228,90,25,193,211,84,239,119,234,85,83,123,120,131,203,219,10,82,35,120,180,249,106,37,169,225,54,103,55,166,124]
    ret = Solution().rob(nums)
    print(ret)
    import timeit
    timer = timeit.Timer("Solution().rob(%s)" % nums, "from __main__ import Solution")
    print("timeConsume: %s" % timer.timeit(10))