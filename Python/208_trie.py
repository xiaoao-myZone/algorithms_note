#/usr/bin/env python
# -*- coding: utf-8 -*-
"""
执行用时：520 ms, 在所有 Python 提交中击败了8.75% 的用户
内存消耗：36.3 MB, 在所有 Python 提交中击败了45.19% 的用户
Trie.next用list

执行用时：404 ms, 在所有 Python 提交中击败了21.28% 的用户
内存消耗：42.2 MB, 在所有 Python 提交中击败了5.26% 的用户
Trie.next用dict
"""

"""
pre:
    1. 什么是前缀, 一个单词本身算不算是自己的前缀
    2. produce和producer,前着即使
    3. 题目没有说清楚如何测试

post:
    1.

"""

class Trie(object):
    def __init__(self, val=None, left=None, right=None):
        """
        Initialize your data structure here.
        """
        self.val = val
        #self.next = []
        self.next = {}
        self.is_end = False
    
    def insert(self, word):
        """
        Inserts a word into the trie.
        :type word: str
        :rtype: None
        """
        if word:
            i = word[0]
            the_next = self.next.get(i)
            if not the_next:
                the_next = Trie(i)
                #self.next.append(the_next)
                self.next[i] = the_next
            the_next.insert(word[1:])
        else:
            if not self.is_end:
                self.is_end = True
    
    def search(self, word):
        """
        Returns if the word is in the trie.
        :type word: str
        :rtype: bool
        """
        if word:
            i = word[0]
            the_next = self.next.get(i)
            if the_next:
                return the_next.search(word[1:])
            # for i in range(len(self.next)):
            #     the_next = self.next[i]
            #     if self.next[i].val == word[0]:
            #         break
            else:
                return False
        else:
            return self.is_end

        # if word:
        #     for i in range(len(self.next)):
        #         the_next = self.next[i]
        #         if self.next[i].val == word[0]:
        #             break
        #     else:
        #         return False
        #     return the_next.search(word[1:])
        # else:
        #     return self.is_end
        
    
    def startsWith(self, prefix):
        """
        Returns if there is any word in the trie that starts with the given prefix.
        :type prefix: str
        :rtype: bool
        """
        if prefix:
            i = prefix[0]
            the_next = self.next.get(i)
            if the_next:
                return the_next.startsWith(prefix[1:])
            else:
                return False
        else:
            return True

# Your Trie object will be instantiated and called as such:
# obj = Trie()
# obj.insert(word)
# param_2 = obj.search(word)
# param_3 = obj.startsWith(prefix)

if __name__ == "__main__":
    obj = Trie()
    for i in ["Trie","insert","search","search","startsWith","insert","search", "apple"]:

        obj.insert(i)
    word = 'app'
    # print(param_2)
    
    param_2 = obj.search(word)
    #from ipdb import set_trace; set_trace()
    param_3 = obj.startsWith(word)
    print(param_2, param_3)
    # from timeit import timeit
    # n = 0
    # ret = Solution()
    # print(ret)
    # import timeit
    # timer = timeit.Timer("Solution().functionCall(%s)" % n, "from __main__ import Solution")
    # print("timeConsume: %s" % timer.timeit(10))