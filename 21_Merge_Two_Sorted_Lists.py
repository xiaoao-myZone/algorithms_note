#/usr/bin/env python
# -*- coding: utf-8 -*-
"""
执行用时：20 ms, 在所有 Python 提交中击败了89.51% 的用户
内存消耗：12.9 MB, 在所有 Python 提交中击败了83.76% 的用户

"""

"""
pre:
    1. 

post:
    1.

"""
# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

class Solution(object):
    def mergeTwoLists(self, l1, l2):
        """将两个升序列表合并为一个升序列表
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        tmp = ListNode()
        ret  = tmp
        while l1 and l2:
            if l1.val>l2.val:
                tmp.next = l2
                l2 = l2.next
            else:
                tmp.next = l1
                l1 = l1.next
            tmp = tmp.next
        
        if l1:
            tmp.next = l1
        else:
            tmp.next = l2
        
        return ret.next
                

if __name__ == "__main__":
    from timeit import timeit
    l1 = ListNode(1)
    tmp = l1
    tmp.next = ListNode(3); tmp = tmp.next
    tmp.next = ListNode(5); tmp = tmp.next
    l2 = ListNode(2)
    tmp = l2
    tmp.next = ListNode(2); tmp = tmp.next
    tmp.next = ListNode(4); tmp = tmp.next
    ret = Solution().mergeTwoLists(l2, l1)
    while ret:
        print(ret.val)
        ret = ret.next
    
    # import timeit
    # timer = timeit.Timer("Solution().mergeTwoLists(l1, l2)", "from __main__ import Solution")
    # print("timeConsume: %s" % timer.timeit(10))