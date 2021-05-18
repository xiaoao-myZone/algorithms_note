
"""
1. 一个源头栈出栈顺序为1，2，3，4，5，这些数值进入到一个缓存栈中
2. 这个缓存栈随时可能向目标栈出栈
3. 最终缓存栈与源头栈均为空，所有数字均在目标栈中

请根据目标栈的数字顺序，判断其是否可能是通过上述过程得到的
"""
from random import choice

# understand
def gain():
    input_stack = [5, 4, 3, 2, 1]
    tmp = []
    out_put = []
    while input_stack or tmp:
        if input_stack:
            tmp.append(input_stack.pop())
        if choice([True,False]):
            if tmp:
                out_put.append(tmp.pop())
            else:
                continue
    return out_put

print(gain())
source_stack = [5, 4, 3, 2, 1] # input: append # output: pop
target_stack = [1, 3, 4, 2, 5]

def solution(source_stack, target_stack):
    msg = []
    temp = "{} | -- {} --> | {}"
    length = len(target_stack)
    if length != len(source_stack):
        return False
    target_stack.reverse()
    tmp = []
    point_s = 0
    point_t = 0
    while point_s<length and point_t<length:
        i = target_stack[point_t]
        j = source_stack[point_s]
        if i == j:
            point_s+=1
            point_t+=1
            msg.append(temp.format("tmplate", i, "target "))
            msg.append(temp.format("source ", i, "tmplate"))
        else:
            if tmp and tmp[-1] == j:
                point_s += 1
                tmp.pop()
                msg.append(temp.format("source ", j, "tmplate "))
            else:
                tmp.append(i)
                point_t += 1
                msg.append(temp.format("tmplate", i, "target "))

    if point_s == len(source_stack):
        for i in msg[::-1]:
            print(i)
        return True
    else:
        return False
#from ipdb import set_trace; set_trace()
res=solution(source_stack, target_stack)
print(res)


def high(source_stack, target_stack):
    length = len(target_stack)
    if length != len(source_stack):
        return False
    target_stack.reverse()
    tmp = []
    point_s = 0
    point_t = 0
    while point_s<length and point_t<length:
        i = target_stack[point_t]
        j = source_stack[point_s]
        if i == j:
            point_s+=1
            point_t+=1
        else:
            if tmp and tmp[-1] == j:
                point_s += 1
                tmp.pop()
            else:
                if tmp:
                    if i<tmp[-1]:
                        return False
                tmp.append(i)
                point_t += 1
    return True

    

print(high(source_stack, target_stack))

"""
Conclusion: 这是不是就是传说中的堆排序?
"""