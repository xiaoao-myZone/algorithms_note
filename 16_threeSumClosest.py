"""
执行用时：2784 ms, 在所有 Python 提交中击败了7.34% 的用户
内存消耗：12.9 MB, 在所有 Python 提交中击败了77.16% 的用户
傪胜

"""

"""
pre:
    1. 先将数组排序为升序
    2. 每级循环中均可以判断趋势为上升, 如果delta>0, 那么就没有必要接着循环了(未必，只能是最后一个内循环)
    3. 是求delta最小值,不知道问题边界的情况下,只能赋上随便三个数与目标值的
    4. 如果δ的值从负值变为正值

post:
    1.还好返回值是和，不需要时具体的三个数， 所以pre.3可以成立

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
        if not delta:
            return target
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
                    if not vary_delta:
                        return target
                    if abs(vary_delta) < abs(delta):
                        delta = vary_delta
                    else: #进入else不可能为0
                        if vary_delta>0:
                            break
                        # if delta<0:
                        #     if vary_delta<0:
                        #         pass
                        #     else:
                        #         pass
                        
                        # else:
                        #     if vary_delta>0:
                        #         pass
                        #     else:
                        #         pass
                        

                    tmp+=1
                j+=1
            i += 1
        return target + delta
if __name__ == "__main__":
    nums = [-4, -1, 1, 2]
    target = 1
    ret = Solution().threeSumClosest(nums, target)
    print(ret)