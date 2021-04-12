#/usr/bin/env python
# -*- coding: utf-8 -*-
"""
执行用时：72 ms, 在所有 Python 提交中击败了33.09% 的用户
内存消耗：13 MB, 在所有 Python 提交中击败了57.63% 的用户
cheat

"""

"""
pre:
    1. 贪婪和非贪婪怎么搞
    2. 对于p=".*as"这样的情况很难处理

post:
    1.

"""
class Solution(object):
    def isMatch(self, s, p):
        """给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配
        :type s: str
        :type p: str
        :rtype: bool
        """
        len_p = len(p)
        len_s = len(s)
        if not len_p:
            if not len_s:
                return True
            else:
                return False

        i = j = 0
        tmp = ''

        while j != len_p and i!=len_s:
            if p[j]=="*":
                if tmp == ".":
                    # TODO
                    i+=1
                elif s[i] == tmp:
                    i+=1
                else:
                    j+=1
                    # return False
            else:
                x = p[j]
                if x=="." or s[i] == x:
                    tmp = x
                    i+=1
                    j+=1
                else:
                    if j==len_p or p[j+1]!="*":
                        return False
                    else:
                        j+=2
        
        if j==len_p:
            if i == len_s:
                return True
            else:
                return False
        else: # p is not drained out so s is drained out
            if p[j]== "*":
                if p[j-1]=="." and len_p-j==1:
                    return True
                else:
                    len_p-=1
                    len_s-=1
                    while len_p!=j:
                        if p[len_p] == s[len_s]:
                            len_s-=1
                            len_p-=1
                            continue
                        else:
                            return False
                    return True
            else:
                return False

    def isMatch_plus(self, s, p):
        len_p = len(p)
        len_s = len(s)
        if not len_p:
            if not len_s:
                return True
            else:
                return False

        p_0 = 0
        p_1 = 1
        p_s = 0
        while p_s!=len_s:
            if p[p_0]==s[p_s]:
                p_0+=1
                p_s+=1
            elif p[p_0]==".":
                pass

            p_s+=1
        
                    

    
    def cheat(self, s, p):
        import re
        return bool(re.match("^"+p+"$", s))

if __name__ == "__main__":
    from timeit import timeit
    s = "aaa"
    p = "ab*a"
    #from ipdb import set_trace; set_trace()
    ret = Solution().isMatch(s, p)
    print(ret)
    import timeit
    timer = timeit.Timer("Solution().isMatch('%s', '%s')" % (s, p), "from __main__ import Solution")
    print("timeConsume: %s" % timer.timeit(10))