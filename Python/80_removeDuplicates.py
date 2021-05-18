"""
执行用时：28 ms, 在所有 Python 提交中击败了38.93% 的用户
内存消耗：13.1 MB, 在所有 Python 提交中击败了31.39% 的用户
"""

"""
pre:
    1. 

post:
    1.

"""
class Solution(object):
    def removeDuplicates(self, nums):
        """使有序数组中每个数字最多只有两两个
        :type nums: List[int]
        :rtype: int
        """
        i = 0
        num = nums[i]
        bear = False
        i+=1
        length = len(nums)
        while i<length:
            if not bear:
                if num == nums[i]:
                    bear = True
                else:
                    num = nums[i]
                i+=1
            else:
                if num == nums[i]:
                    del nums[i]
                    length-=1
                else:
                    num = nums[i]
                    i+=1
                    bear = False
        return length
    



if __name__ == "__main__":
    nums = [1,1,1,2,3,3,5,5,5,6,6,6]
    ret = Solution().removeDuplicates(nums)
    print(ret)
    print(nums)



