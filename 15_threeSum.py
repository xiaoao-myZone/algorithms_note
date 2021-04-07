"""
执行用时：3176 ms, 在所有 Python 提交中击败了9.53% 的用户
内存消耗：18.6 MB, 在所有 Python 提交中击败了33.35% 的用户
惨胜
Post:
    1. 将数组排序后分为两端(为0的区段可以重合),再对这两端进行排列组合的结果会有重复
    2. 1的结果不会出现重复,因为所有的结果是按顺序排列的

运行时间超时qvq
"""

class Solution(object):
    def threeSum(self, nums):
        """
        :type nums: List[int]
        :rtype: List[List[int]]
        """
        nums.sort()
        length = len(nums)
        ret = []
        lPoint = 0
        rPoint = length - 1
        tmp = 0
        for lPoint in range(length):
            print("left point: %d | value: %d" % (lPoint, nums[lPoint]))
            if nums[lPoint]>0:
                break
            if lPoint:
                if nums[lPoint] == nums[lPoint-1]:
                    continue
            
            for rPoint in range(length-1, 0, -1):
                # if nums[lPoint] == -70 and nums[rPoint] == 76:
                #     from ipdb import set_trace; set_trace()
                print("   right point: %d | value: %d" % (rPoint, nums[rPoint]))
                if lPoint == rPoint:
                    print("   equal")
                    break
                if nums[rPoint]<0:
                    print("   minus")
                    break
                if rPoint<length-1:
                    if nums[rPoint] == nums[rPoint+1]:
                        print("   step same num")
                        continue
                tmp = -(nums[lPoint] + nums[rPoint])
                # for i in range(lPoint+1, rPoint):
                #     print("      middle point: %d | value: %d" % (i, nums[i]))
                #     if tmp == nums[i]:
                i = self.find_res(nums, lPoint, rPoint, tmp)
                if i is not None:
                    ret.append([nums[lPoint], nums[i], nums[rPoint]])
        return ret
    
    def find_res(self, nums, start, end, target):
        while (end-start)>1:
            middle = (start + end)//2
            if nums[middle]<target:
                start = middle
            elif nums[middle]>target:
                end = middle
            else:
                return middle
        return None


if __name__ == "__main__":
    #nums = [-1,0,1,2,-1,-4]
    nums = [-1,0,0,0,0,-4]
    nums = [34,55,79,28,46,33,2,48,31,-3,84,71,52,-3,93,15,21,-43,57,-6,86,56,94,74,83,-14,28,-66,46,-49,62,-11,43,65,77,12,47,61,26,1,13,29,55,-82,76,26,15,-29,36,-29,10,-70,69,17,49]
    # 43
    ret = Solution().threeSum(nums)
    print(ret)
    print(len(ret))
    print(nums)
    # nums.sort()
    # print(nums)
    # ret = Solution().find_res(nums, 0, 6, 0)
    # print(ret)

"""
[[-82,-11,93],[-82,13,69],[-82,17,65],[-82,21,61],[-82,26,56],[-82,33,49],[-82,34,48],[-82,36,46],[-70,-14,84],[-70,-6,76],[-70,1,69],[-70,13,57],[-70,15,55],[-70,21,49],[-70,34,36],[-66,-11,77],[-66,-3,69],[-66,1,65],[-66,10,56],[-66,17,49],[-49,-6,55],[-49,-3,52],[-49,1,48],[-49,2,47],[-49,13,36],[-49,15,34],[-49,21,28],[-43,-14,57],[-43,-6,49],[-43,-3,46],[-43,10,33],[-43,12,31],[-43,15,28],[-43,17,26],[-29,-14,43],[-29,1,28],[-29,12,17],[-29,-14,43],[-29,1,28],[-29,12,17],[-14,-3,17],[-14,1,13],[-14,2,12],[-11,-6,17],[-11,1,10],[-3,1,2]]

[[-82,-11,93],[-82,13,69],[-82,17,65],[-82,21,61],[-82,26,56],[-82,33,49],[-82,34,48],[-82,36,46],[-70,-14,84],[-70,-6,76],[-70,1,69],[-70,13,57],[-70,15,55],[-70,21,49],[-70,34,36],[-66,-11,77],[-66,-3,69],[-66,1,65],[-66,10,56],[-66,17,49],[-49,-6,55],[-49,-3,52],[-49,1,48],[-49,2,47],[-49,13,36],[-49,15,34],[-49,21,28],[-43,-14,57],[-43,-6,49],[-43,-3,46],[-43,10,33],[-43,12,31],[-43,15,28],[-43,17,26],[-29,-14,43],[-29,1,28],[-29,12,17],[-14,-3,17],[-14,1,13],[-14,2,12],[-11,-6,17],[-11,1,10],[-3,1,2]]

"""