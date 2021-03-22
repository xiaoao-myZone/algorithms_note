#include <stdio.h>
#include <malloc.h>
#include <stdbool.h>
#include <string.h>

/*
执行用时：4 ms, 在所有 C 提交中击败了95.72% 的用户
内存消耗：5.7 MB, 在所有 C 提交中击败了75.55% 的用户
*/

/*
pre:
    1. 从哪边遍历输入字符串?从小到大,处理起来比较高效与方便,只是需要逆序输出
    2. 谨慎使用for,有时候状态不需要改变,没有while灵活

post:
    1.

*/
int romanToInt(char * s){
    char * sign = "IVXLCDM";
    char len = strlen(s)-1;
    s+=len;
    int res = 0;
    short bit = 1;
    bool min=false;
    for(;len>-1;len--,s--){
        printf("%s\n", sign);
        if (*s==sign[0]){
            if (min) {
                res -= bit*1;
                bit*=10;
                sign+=2;
            }
            else res += bit*1;
            min = false; //由于5和9这种情况,地位在高位前只会出现一次
            
        }
        else if (*s==sign[1]){
            res += bit*5;
            min = true;
        }
        else if (*s==sign[2]){
            printf("five\n");
            res += bit*10;
            min = true;
        }
        else {
            bit*=10;
            sign+=2;
            min = false;
            s++;
            len++;
            printf("next\n");
        }

    }
    return res;
}

void main(void){
    char * s = "LVIII";
    printf("%d\n", romanToInt(s));
    
}
