#/usr/bin/env python
# -*- coding: utf-8 -*-
"""

"""

"""
pre:
    1. 

post:
    1.

"""
class Solution(object):
    def removeDuplicates(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        i = 0
        j = 1
        length = len(nums)
        if length<2:
            return length
        while j<length:
            if nums[i] == nums[j]:
                del nums[j]
                length-=1
            else:
                i+=1
                j+=1

        return length

if __name__ == "__main__":
    from timeit import timeit
    n = [0,0,1,1,1,2,2,3,3,4]
    ret = Solution().removeDuplicates(n)
    print(ret)
    import timeit
    timer = timeit.Timer("Solution().removeDuplicates(%s)" % n, "from __main__ import Solution")
    print("timeConsume: %s" % timer.timeit(10))