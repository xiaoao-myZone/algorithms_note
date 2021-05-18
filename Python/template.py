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