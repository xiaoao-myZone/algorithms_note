"""
执行用时：24 ms, 在所有 Python 提交中击败了43.14% 的用户
内存消耗：13.3 MB, 在所有 Python 提交中击败了45.66% 的用户

"""

"""
pre:
    1. 与153相同
    2. leetcode有灌水嫌疑

post:
    1.

"""
class Solution(object):
    def findMin(self, nums):
        """将尾部元素放到数组首部，称为一次旋转，nums是经过若干次旋转后得到的，原数组为升序，且可能有相同值，要求返回最小值
        :type nums: List[int]
        :rtype: int
        """
        tmp = nums[0]
        for i in nums[1:]:
            if i<tmp:
                return i
            else:
                tmp = i
        return nums[0]

if __name__ == "__main__":
    nums = [2,2,2,0,1]
    ret = Solution().findMin(nums)
    print(ret)