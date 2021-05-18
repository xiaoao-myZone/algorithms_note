#/usr/bin/env python
# -*- coding: utf-8 -*-
"""
执行用时：12 ms, 在所有 Python 提交中击败了98.45% 的用户
内存消耗：12.9 MB, 在所有 Python 提交中击败了95.35% 的用户

"""

"""
pre:
    1. 

post:
    1.

"""
#Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution(object):
    def minDiffInBST(self, root):
        """
        :type root: TreeNode
        :rtype: int
        """
        res = []
        self.traverse(root, res)
        print(res)
        res.sort()
        ret = res[1] - res[0]
        length = len(res)
        if length>2:
            for i in range(1, length-1):
                tmp = res[i+1]-res[i]
                if ret>tmp:
                    ret = tmp
        return ret
    
    def traverse(self, root, res):
        if root:
            if root.val:
                res.append(root.val)
            self.traverse(root.left, res)
            self.traverse(root.right, res)

    def traverse_create(self, root, res):
        if not res:
            return 
        root.val = res.pop(0)
        if not root.left:
            root.left = TreeNode()
        self.traverse_create(root.left, res)
        if not root.right:
            root.right = TreeNode()
        self.traverse_create(root.right, res)


if __name__ == "__main__":
    from timeit import timeit
    n = [1,3,5,8,7]
    root = TreeNode()
    Solution().traverse_create(root, n)
    ret = Solution().minDiffInBST(root)
    print(ret)
    # import timeit
    # timer = timeit.Timer("Solution().functionCall(%s)" % n, "from __main__ import Solution")
    # print("timeConsume: %s" % timer.timeit(10))