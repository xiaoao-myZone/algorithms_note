#include <string.h>
#include <stdio.h>
#include <malloc.h>

/* 
执行用时：16 ms, 在所有 C 提交中击败了84.44% 的用户
内存消耗：5.9 MB, 在所有 C 提交中击败了84.92% 的用户
*/

/* (取消max的使用后)
执行用时：20 ms, 在所有 C 提交中击败了78.16% 的用户
内存消耗：5.8 MB, 在所有 C 提交中击败了98.44%        
*/

/*最长回文字
[pre]
断论:
    1. 首尾必相等(长度可以为一)
    2. 最长回文里面一定有子回文,或者说找到对称中心,向两边发展,找到最大边界
一般性策略:
    1. 找到aa与aba模式的核(优化:左侧指针先向左移动,开始检查,复位后,左右侧指针再向两边一起移动)
    2. 向外扩张找到最大值,并记录比较
    3. 如何记录结果(最少需要两个值)

[post]
1. 既可以使用从中心(遍历2*n-1个中心)出发的思路,也可以从边缘向中心逼近的思路
2. 但是每次从边缘向中心对比的过程中,不可能不通过新增变量,使得检查完这个边缘后可以返回检查前的状态再遍历下一个边缘
3. 我直觉认为遍历过程中产生的结果可以对提前判断之后的中心是否存在更大的回文字

*/

char * longestPalindrome(char * s){
    int str_len = strlen(s);
    if(str_len==1) return s;
    //printf("strlen: %d\n", str_len);
    

    //int max_len = 0; // I known it should be 1
    int i,j;
    int ret_a, ret_b;
    ret_a = ret_b = 0;
    i = j = 1;

    while (i<str_len){
        i--;
        printf("seg_1--start--i: %d --j: %d\n", i, j);
        while (i>-1 && j<str_len){
            printf("##index: %d: %d\n", i, j);
            if (s[i]==s[j]){
                
                i--;j++;
            }
            else break;
        }
        i++;j--;
        printf("seg_1--end  --i: %d --j: %d\n", i, j);
        if (ret_b-ret_a<j-i){
            //max_len = j-i;
            ret_b = j;
            ret_a = i;
            printf("ret change--a: %d --b: %d\n", ret_a, ret_b);
        }

        i = j = (i+j)/2 + 1; //return back
        i--; j++;
        printf("seg_2--start--i: %d --j: %d\n", i, j);
        while (i>-1 && j<str_len){
            printf("##index: %d: %d\n", i, j);
            if (s[i]==s[j]){
                i--;j++;
            }
            else break;
        }
        i++;j--;
        printf("seg_2--end  --i: %d --j: %d\n", i, j);
        if (ret_b-ret_a<j-i){
            //max_len = j-i;
            ret_b = j;
            ret_a = i;
            printf("ret change--a: %d --b: %d\n", ret_a, ret_b);
        }
        i = j = (i+j)/2 + 1;
    }
    
    printf("a:%d, b:%d\n", ret_a, ret_b);
    s[ret_b+1] = '\0';
    s+=ret_a;
    
    return s;
}

void main(void){
    char s[20] = "ccc\0";
    //printf("str_len: %ld\n", strlen(s));
    printf("|%s|\n", longestPalindrome(s));
}
