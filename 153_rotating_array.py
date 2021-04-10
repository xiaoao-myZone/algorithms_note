"""
执行用时：16 ms, 在所有 Python 提交中击败了84.88% 的用户
内存消耗：13 MB, 在所有 Python 提交中击败了82.80% 的用户
感觉在灌水
"""

"""
pre:
    1. 这意味着前半段为降序，后半段为升序
    2. 只需要找到升/降序的转折点，时间复杂度应为O(n)

post:
    1.

"""
class Solution(object):
    def findMin(self, nums):
        """将尾部元素放到数组首部，称为一次旋转，nums是经过若干次旋转后得到的，原数组为升序，且无相同值，要求返回最小值
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
    nums=  [1,2,3,4,5]
    ret = Solution().findMin(nums)
    print(ret)