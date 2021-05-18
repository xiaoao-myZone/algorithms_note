#/usr/binum/env python
# -*- coding: utf-8 -*-
"""
执行用时：16 ms, 在所有 Python 提交中击败了97.54% 的用户
内存消耗：13 MB, 在所有 Python 提交中击败了35.09% 的用户
isUgly_3

"""

"""
pre:
    1. 能被2,3,5除为1数即为丑数

post:
    1. 超时的时候未必是算法问题, 看一下输出值, 可能意想不到
    2. 利用以2,5为因数的数的个位数特征, 将三次原数求余简化为两次

"""
class Solution(object):
    def isUgly(self, num):
        """丑数 就是只包含质因数 2、3 和/或 5 的正整数。1 通常被视为丑数。
        :type n: int
        :rtype: bool
        """
        if num is 1:
            return True
        
        if not num%2:
            return self.isUgly(num//2)
        elif not num%3:
            return self.isUgly(num//3)
        elif not num%5:
            return self.isUgly(num//5)

        else:
            return False
        
    def isUgly_1(self, num):
        """丑数 就是只包含质因数 2、3 和/或 5 的正整数。1 通常被视为丑数。
        :type n: int
        :rtype: bool
        """
        while num is not 1:
            for i in (2,3,5):
                if not n%i:
                    num //= i
                    break
            else:
                return False

        return True

    def isUgly_2(self, num):
        """丑数 就是只包含质因数 2、3 和/或 5 的正整数。1 通常被视为丑数。
        :type n: int
        :rtype: bool
        """
        if num is 0:
            return False
        while num is not 1:
            nums = self.split_num(num)
            if not nums[0] % 2:
                num //= 2
                continue
            elif not sum(nums) % 3:
                num //= 3
                continue
            elif nums[0] is 5:
                num //= 5
                continue
            else:
                return False
        return True
            

    
    def split_num(self, num):
        ret = []
        while num>10:
            ret.append(num%10)
            num //= 10
        
        ret.append(num)
        return ret

    def isUgly_3(self, num):
        """丑数 就是只包含质因数 2、3 和/或 5 的正整数。1 通常被视为丑数。
        :type n: int
        :rtype: bool
        """
        if num is 0:
            return False
        while num is not 1:
            last = num % 10
            if last is 5:
                num //= 5
                continue
            elif not last % 2: # last in [2,4,6,8,0]:
                num //= 2
                continue
            elif not num % 3:
                num //= 3
                continue
            else:
                return False

        return True


    
        
        
        

if __name__ == "__main__":
    #from ipdb import set_trace; set_trace()
    n = 14
    ret = Solution().isUgly_2(n)
    print(ret)