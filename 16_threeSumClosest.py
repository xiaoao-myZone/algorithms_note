"""


"""

"""
pre:
    1. 先将数组排序为升序
    2. 每级循环中均可以判断趋势为上升, 如果delta>0, 那么就没有必要接着循环了
    3. 是求delta最小值,不知道问题边界的情况下,只能赋上随便三个数与目标值的差

post:
    1.

"""
class Solution(object):
    def threeSumClosest(self, nums, target):
        """给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，
           使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
           
        :type nums: List[int]
        :type target: int
        :rtype: int
        """
        nums.sort()
        delta = sum(nums[:3]) - target
        vary_delta = 0
        length = len(nums)
        margin_i = length-2
        margin_j = margin_i+1
        i = 0
        j = 2
        while i<margin_i:
            tmp_i = nums[i]
            j = i+1
            while j<margin_j:
                tmp_j = nums[j]
                tmp = j+1
                while tmp<length:
                    vary_delta = tmp_i+tmp_j+nums[tmp] - target
                    if abs(vary_delta) < abs(delta):
                        delta = vary_delta
                    else:
                        if delta<0:
                            if vary_delta<0:
                                pass
                            else:
                                pass
                        
                        else:
                            if vary_delta>0:
                                pass
                            else:
                                pass
                        

                    tmp+=1
                j+=1

            i += 1
if __name__ == "__main__":
    ret = Solution()