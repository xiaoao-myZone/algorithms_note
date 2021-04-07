"""
执行用时：24 ms, 在所有 Python 提交中击败了53.17% 的用户
内存消耗：13.2 MB, 在所有 Python 提交中击败了81.22% 的用户

"""

"""
pre:
    1. 可以将数组还原后再用二分法搜索 n + log(n)

post:
    1. 使用while 循环的时候先想好如何改变循环判断状态

"""
class Solution(object):
    def search(self, nums, target):
        """一个非降序数组,从某处打断,将换位置再拼接后得到新数组,判断某一个数是否在其中
        :type nums: List[int]
        :type target: int
        :rtype: bool
        """
        length = len(nums)
        i = length - 1
        while i:
            if nums[i] < nums[i-1]:
                nums = nums[i:] + nums[:i]
                break
            i-=1

        a = -1
        b = length
        i = 0
        while b-a>1:
            i = (a+b)//2
            if nums[i]<target:
                a = i
            elif nums[i]>target:
                b = i
            else:
                return True
        return False

        



if __name__ == "__main__":
    nums = [2,5,6,0,0,1,2]
    target = 5
    ret = Solution().search(nums, target)
    print(ret)