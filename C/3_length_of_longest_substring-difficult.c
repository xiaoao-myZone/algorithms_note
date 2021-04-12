#include <stdio.h>
#include <string.h>
/* 85 17 */

/*

0 <= s.length <= 5 * 104
s 由英文字母、数字、符号和空格组成

*/
int old_lengthOfLongestSubstring(char * s){
    if (!s[0]) return 0;
    int a,b;
    int res=0;
    a=b=0;
    int i=0;
    while (s[++i])
    {
        b = i;
        for (int j=a; j<i; j++)
        {
            if (s[j] == s[i])
            {
                if (b-a> res) // b-(a+1)-1
                {
                    res = b-a;
                }
                a = j+1;
                break;

            }
        }
    }
    if (b-a+1> res)
    {
        res = b-a+1;
    }
    return res;
}
/*
1. 这个函数的复杂度必然是n^2
2. 这个函数需要声明一个整形来存放结果,并且必须不断和新结果比较,取最大值
3. 这个函数必须声明两个整形来标记起点和终点
*/
int lengthOfLongestSubstring(char * s){
    int i;
    int res;
    i = res = 0;
    for (int j=0; j<strlen(s);j++){

    }
}

void main(void)
{
    char a[] = "aaab";
    int res=lengthOfLongestSubstring(a);
    printf("\nmaxlen is %d", res);

    

}