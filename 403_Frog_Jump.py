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
    pass

if __name__ == "__main__":
    from timeit import timeit
    n = 0
    ret = Solution()
    print(ret)
    import timeit
    timer = timeit.Timer("Solution().functionCall(%s)" % n, "from __main__ import Solution")
    print("timeConsume: %s" % timer.timeit(10))




"""
一只青蛙想要过河。 假定河流被等分为若干个单元格，并且在每一个单元格内都有可能放有一块石子（也有可能没有）。 青蛙可以跳上石子，但是不可以跳入水中。

给你石子的位置列表 stones（用单元格序号 升序 表示）， 请判定青蛙能否成功过河（即能否在最后一步跳至最后一块石子上）。

开始时， 青蛙默认已站在第一块石子上，并可以假定它第一步只能跳跃一个单位（即只能从单元格 1 跳至单元格 2 ）。

如果青蛙上一步跳跃了 k 个单位，那么它接下来的跳跃距离只能选择为 k - 1、k 或 k + 1 个单位。 另请注意，青蛙只能向前方（终点的方向）跳跃。

 

示例 1：

输入：stones = [0,1,3,5,6,8,12,17]
输出：true
解释：青蛙可以成功过河，按照如下方案跳跃：
跳 1 个单位到第 2 块石子
然后跳 2 个单位到第 3 块石子
接着 跳 2 个单位到第 4 块石子
然后跳 3 个单位到第 6 块石子
跳 4 个单位到第 7 块石子
最后，跳 5 个单位到第 8 个石子（即最后一块石子）。

"""