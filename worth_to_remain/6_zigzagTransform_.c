#include <stdio.h>
#include <string.h>
#include <malloc.h>
/*
执行用时：8 ms, 在所有 C 提交中击败了80.71% 的用户
内存消耗：6.3 MB, 在所有 C 提交中击败了67.78% 的用户
*/

/*
Pre:
    1. 直观的做法是将每一行排好后,再将各行依次串起
    2. 可以创建若干子串,然后将其并起, 但是串不是链表,拼接需要消耗额外的内存
    3. 最好只创建一个串,然后计算每行的初始位置,再通过循环将字母依次填入
    4. 得出通用公式: 2*l*(RowNum-1)+r与2*(l+1)*(RowNum-1)-r 其中l是列数(从0开始), r是行数(从0开始)
    5. 从数学上来讲,就是把一个有序序列,变换成另一种序列,一一映射
    6. 如果让原串递增,那么目标串需要需要计算并保存每行的最大值,在每行的初始位置增加偏移量(列数)
    7. 如果让目标串递增,那么原串需要根据公式跳跃

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
    char * ret = (char *)malloc(sizeof(char)*(length+1));
    char * tmp_ret = ret;
    char temp = 0;
    ret[length] = '\0';
    // first line
    //printf("first: %d\n", (length-1)/(2*(rowNum-1))+1);
    for(int i=0;i<(length-1)/(2*(rowNum-1))+1;i++){
        
        ret[0] = s[2*i*(rowNum-1)];
        
        //printf("index: %d, value: %c\n", 2*i*(rowNum-1), ret[0]);
        ret++;
    }
    for(int j=1; j<rowNum-1;j++){
        //(length-1-j)/(2*(rowNum-1))+1
        for(int i=0;;i++){
            temp = 2*i*(rowNum-1)+j;
            if (temp<length){
                ret[0] = s[temp];
                ret++;
            }
            else break;//同一层,偶数(从0开始)列一定不小于奇数列
            temp = 2*(i+1)*(rowNum-1)-j;
            if (temp<length){
                ret[0] = s[temp];
                ret++;
            }

        }
    }


    //last line
    for(int i=0;i<(length-rowNum)/(2*(rowNum-1))+1;i++){
        
        ret[0] = s[(2*i+1)*(rowNum-1)];
        
        //printf("index: %d, value: %c\n", 2*i*(rowNum-1), ret[0]);
        ret++;
    }
    return tmp_ret;

}

void main(void)
{
    //char *s = "xiaoaoxiaoyi";
    char *s = "A";//"PAYPALISHIRING"-->"PAHNAPLSIIGYIR"
    // char * r = convert_2(s, 3);
    // for (int i=0;i<strlen(s);i++)
    // printf("%c", r[i]);
    // printf("\n%s", r);
    printf("|%s|\n",convert(s, 3));

}