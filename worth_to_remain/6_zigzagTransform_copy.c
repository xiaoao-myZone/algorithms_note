#include <stdio.h>
#include <string.h>
#include <malloc.h>
/*
执行用时：8 ms, 在所有 C 提交中击败了80.71% 的用户
内存消耗：6.3 MB, 在所有 C 提交中击败了67.78% 的用户
*/
/*
执行用时：4 ms, 在所有 C 提交中击败了95.19% 的用户
内存消耗：6.3 MB, 在所有 C 提交中击败了85.56% 的用户   -*- 第二刷 -*-
*/

/*
Pre:
    1. 直观的做法是将每一行排好后,再将各行依次串起
    2. 可以创建若干子串,然后将其并起, 但是串不是链表,拼接需要消耗大量额外的内存
    3. 最好只创建一个串,然后计算每行的初始位置,再通过循环将字母依次填入
    4. 得出通用公式: 2*l*(RowNum-1)+r与2*(l+1)*(RowNum-1)-r 其中l是列数(从0开始), r是行数(从0开始)
    5. 从数学上来讲,就是把一个有序序列,变换成另一种序列,一一映射
    6. 如果让原串递增,那么目标串需要来回折返各个层,这需要计算并保存每行的最大值,在每行的初始位置增加偏移量(列数),这样空间消耗就很大
    7. 如果让目标串递增,那么原串的索引需要根据公式跳跃

*/

/*
Post:
    1. 本质上是一个数学题
    2. 中间任意一层,偶数列一定不小于奇数列,这个判断节省了很多行代码
*/

char * convert_2(char * s, int rowNum)
{
    if (rowNum==1) return s;
    int len = strlen(s)+1;
    printf("\nlen is %d\n", len);
    char * ret = (char * )malloc(sizeof(char)*len);
    int i=0;
    int j=0;
    int col = 0;
    int row=0;

    for (int col=0;;col++,i++)
    {
        j = 2*col * (rowNum-1);
        if (j>len-2){
            printf("give up: %d\n", j);
            break;
        };
        ret[i] = s[j];
        printf("s[%d]: %c --> ret[%d]\n", j,s[j],i);
    }

    for (row=1;row<rowNum-1;row++)
    {
        for (int col=0;;col++,i++)
        {
            if (col % 2)
            {
                j = col * (rowNum-1) + rowNum-1-row;
            }
            else
            {
                j = col * (rowNum-1) + row;
            }
            if (j>len-2){
                printf("give up: %d\n", j);
                break;
            };
            ret[i] = s[j];
            printf("s[%d]: %c --> ret[%d]\n", j,s[j],i);
        }


    }
    row = rowNum-1;
    for (int col=0;;col++,i++)
    {
        j = (2*col + 1) * (rowNum-1);
        if (j>len-2){
            printf("give up: %d\n", j);
            break;
        };
        ret[i] = s[j];
        printf("s[%d]: %c --> ret[%d]\n", j,s[j],i);
    }
    
    ret[len-1] = '\0';
    return ret;

}

char * convert(char * s, int rowNum)
{
    if (rowNum==1) return s;
    int length = strlen(s);
    //printf("\nlen is %d\n", length);
    if (length<=rowNum) return s;
    rowNum -= 1;
    char * ret = (char *)malloc(sizeof(char)*(length+1));
    char * tmp_ret = ret;
    int temp = 0;
    ret[length] = '\0';

    // first line
    //printf("first: %d\n", (length-1)/(2*rowNum)+1);
    for(int i=0;i<(length-1)/(2*rowNum)+1;i++){
        
        ret[0] = s[2*i*rowNum];
        
        //printf("index: %d, value: %c\n", 2*i*rowNum, ret[0]);
        ret++;
    }
    for(int j=1; j<rowNum;j++){//j 为row, i为col
        for(int i=0;;i++){
            temp = 2*i*rowNum+j;
            if (temp<length){
                ret[0] = s[temp];
                //printf("index: %d, value: %c, row: %d, col: %d\n", temp, ret[0], j, i);
                ret++;
            }
            else break;//同一层,偶数(从0开始)列数目一定不小于奇数列
            temp = 2*(i+1)*rowNum-j;
            if (temp<length){
                ret[0] = s[temp];
                //printf("index: %d, value: %c, row: %d, col: %d\n", temp, ret[0], j, i);
                ret++;
            }
            //if(temp<0) return s;

        }
    }


    //last line
    printf("*********%d*********", (length-rowNum)/(2*(rowNum-1))+1);
    for(int i=0;i<(length-rowNum-1)/(2*rowNum)+1;i++){
        
        ret[0] = s[(2*i+1)*rowNum];
        
        //printf("index: %d, value: %c\n", 2*(i+1)*rowNum, ret[0]);
        ret++;
    }
    return tmp_ret;

}

void main(void)
{
    //char *s = "xiaoaoxiaoyi";
    //char *s = "Apalindromeisaword,phrase,number,orothersequenceofunitsthatcanbereadthesamewayineitherdirection,withgeneralallowancesforadjustmentstopunctuationandworddividers.";
    char *s = "PAYPALISHIRING";

    //"PAYPALISHIRING"-->"PAHNAPLSIIGYIR"
    printf("|%s|\n",convert(s, 4));

}