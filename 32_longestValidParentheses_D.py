"""
执行用时：24 ms, 在所有 Python 提交中击败了98.45% 的用户
内存消耗：12.9 MB, 在所有 Python 提交中击败了99.86% 的用户

"""

"""
pre:
    1. 栈
    2. 无效无非是这么几种情况： 
        ①有头没尾(左括号多)，最后才能检查到，最大有效长度可以在余栈中减掉 #也是循环结束的情况 
        ②有尾没头(右括号多)，一出现就可以检查到，对比缓存最大长度即可

post:
    1. 如果有头没尾出现在中间，且之前已经消掉了几对括号， 需要将
    2. 未必一定要用栈
    3. 可以双指针, 不过会重复那一元素和其他元素对比
    4. 如果坚持用栈, 也可以, 需要找到连续出栈, 栈空
    5. 对于"((()()(((())()()"这样的情况,最后的queue里只有(((, 最终结算时, 而并不知道之前的结算在哪里为止
    6. 对此,不仅需要记录中途的消除括号数,还需要记录它是现存的第几个括号

    7. 成功
    8. 这里只有'(', ')'两个值,不需要专门用栈存放, 如果不是这个题的特殊性, 甚至不需要栈, 一个数值来计数就可以
    9. 最后还是用到的了栈, 是因为最终的结果需要记录, 并且将聚焦点放在现存左括号数量上
    10. 与plus相比, 第一次成功的答案没有很准确理解栈的定义--当所有左括号都消失时, log内应有值,就是子结果,类似与函数调用的main的地位
    11. 无需用sum(log[left_bracket:]), 因为栈的回退是一层一层的, 不会出现跃层

"""
class Solution(object):
    def longestValidParentheses(self, s):
        """给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
        :type s: str
        :rtype: int
        """
        left_brackets = 0
        tmp = 0
        log = []
        for i in s:
            #print("***%s" % log)
            if i == '(':
                left_brackets += 1
                log.append(0)
                print("add %s" % log)
            else:
                if left_brackets:
                    left_brackets -= 1
                    #print("log: %s, left_brackets: %d" % (log, left_brackets))
                    log[left_brackets] = sum(log[left_brackets:]) + 2
                    print("log: %s, left_brackets: %d" % (log, left_brackets))
                    del log[left_brackets+1:]

                else:
                    if log:
                        max_len = max(log)
                        tmp = tmp if tmp>max_len else max_len
                        log = [0]
                        left_brackets = 0
                    
        print(log)
        if log:
            max_len = max(log)
            tmp = tmp if tmp>max_len else max_len
        return tmp

    def longestValidParentheses_plus(self, s):
        """给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
        :type s: str
        :rtype: int
        """
        left_brackets = 0
        tmp = 0
        log = [0]
        for i in s:
            #print("***%s" % log)
            if i == '(':
                left_brackets += 1
                log.append(0)
                print("add %s" % log)
            else:
                if left_brackets:
                    left_brackets -= 1
                    #log[-1] = log.pop() + 2 + log[-1]
                    log[-1] += log.pop() + 2
                    print("log: %s, left_brackets: %d" % (log, left_brackets))

                else:

                    max_len = max(log)
                    tmp = tmp if tmp>max_len else max_len
                    log = [0]
                    left_brackets = 0
                    
        print(log)
        if log:
            max_len = max(log)
            tmp = tmp if tmp>max_len else max_len
        return tmp


if __name__ == "__main__":
    #from ipdb import set_trace; set_trace()
    #s= "()))((((()" #"(()(()" 
    #s= "(()(()" 
    # s = "((()))"
    s=")()()(()())"
    ret = Solution().longestValidParentheses(s)
    print()
    print(ret)