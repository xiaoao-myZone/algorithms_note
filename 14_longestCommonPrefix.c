#include <stdio.h>
#include <malloc.h>
#include <string.h>

/*
执行用时：4 ms, 在所有 C 提交中击败了69.38% 的用户
内存消耗：5.9 MB, 在所有 C 提交中击败了30.81% 的用户
*/

/*
pre:

post:
    1. 主要卡在main中的测试部分
    2. 测试用空数组阴我

*/



char * longestCommonPrefix(char ** strs, int strsSize){
    if (!strsSize) return "";
    char point = 0;
    while (1){
        if (strs[0][point]=='\0') return strs[0];
        for(char i=1; i<strsSize;i++){
            if (strs[i][point]=='\0'){
                return strs[i];
            }
            if (strs[0][point]!=strs[i][point]){
                strs[0][point]='\0';
                return strs[0];
            }
        } 
        point++;
    }

}

void main(void){
    //char  m[3][20] = {"flower","flow","flight"};
    char ** strs = (char **)malloc(sizeof(char *)*3);
    for(int i; i<3; i++){
        strs[i] = (char *)malloc(sizeof(char )*20);
    }
    char*m = "flower";
    strcpy(strs[0],m);
    m = "flow";
    strcpy(strs[1],m);
    m = "flight";
    strcpy(strs[2],m);

    printf("res: %s\n", longestCommonPrefix(strs, 3));

}