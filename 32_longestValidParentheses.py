"""

"""

"""
pre:
    1. 栈
    2. 无效无非是这么几种情况： 
        ①有头没尾，最后才能检查到，最大有效长度可以在余栈中减掉 #也是循环结束的情况
        ②有尾没头，一出现就可以检查到，对比缓存最大长度即可

post:
    1. 如果有头没尾出现在中间，且之前已经消掉了几对括号， 需要将
    2. 未必一定要用栈

"""
from queue import LifoQueue
class Solution(object):
    def longestValidParentheses(self, s):
        """给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
        :type s: str
        :rtype: int
        """
        #from ipdb import set_trace; set_trace()
        queue = []
        ret = 0
        tmp = 0
        for i in s:
            if i == "(":
                queue.append(i)
                
            else: # ')' 
                if not queue: #有尾没头
                    if ret<tmp:
                        ret = tmp
        
                    tmp=0
                    continue
                else:
                    queue.pop() #弹出来的一定是'('?是，因为'('没有入栈过
                    tmp+=2
                # else:
                #     if queue.pop()!="(": #有头没尾
                #         ret = tmp
                #     tmp=0
                #     continue
            
        return ret if ret>tmp else tmp


if __name__ == "__main__":
    s= "()(()" #"(()(()"
    ret = Solution().longestValidParentheses(s)
    print(ret)