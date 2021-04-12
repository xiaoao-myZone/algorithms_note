#include <stdio.h>
#include <malloc.h>

/*
执行用时：4 ms, 在所有 C 提交中击败了90% 的用户
内存消耗：5.9 MB, 在所有 C 提交中击败了47% 的用户
*/

/*
pre:
    1. 从小到大,还是从大到小,应该只能从小到大
    2. 如果从小到大,就会逆序将结果输入字符串,难道要用链表?链表的理论长度会是原来的3倍
    3. 再申请一个相同长度的字符串呢
    4. 可以申请一个理论最大长度
    5. 提示1 <= num <= 3999,那么最大长度就是3999,即MMMCMXCIX,长度为9
    6. 因为int最大可知,并且不算很大,所以可以由从高位到低位的计算方式

post:
    1.取一个十进制数的某一位数, num%i/(i/10) i=10*n (n=1,2,3,...)
    2. 本来以为数据长度未知,这样,字符串结果需要申请的变量就很难搞定,看到num的输入范围,就可以知道最大长度了
    3. 如果向优化存储空间,可以将num分段预设字符串长度
    4. 这题耗的时间长,主要原因是如何从高位到低位获取一个10进制数的各个位的数字

*/

/*
字符          数值
I             1
V             5
X             10
L             50
C             100
D             500
M             1000

4开头的数和9开头的数需要特别表达

*/
// typedef struct node {
//     char a;
//     node * next;
// } node;

char * intToRoman(int num){
    char * res = (char *)malloc(sizeof(char)*16);
    char * tmp = res;
    int remainder;
    char * sign = "IVXLCDM";
    //char * sign = "MDCLXVI";
    sign+=6;
    for (int i=10000;i!=1;i/=10,sign-=2){
        printf("i is %d\n", i);
        printf("sign is %s\n", sign);
        remainder = num % i/(i/10);
        printf("reminder is %d\n", remainder);
        if (remainder){
            if (remainder<4){
                while(remainder){
                    //tmp[0] = 'I';
                    tmp[0] = sign[0];
                    tmp++;
                    remainder--;
                }
            }
            else if (remainder==4)
            {
                //tmp[0] = 'V';
                
                //tmp[1] = 'I';
                tmp[1] = sign[1];
                tmp[0] = sign[0];
                tmp+=2;
            }
            else if (remainder==5){
                tmp[0] = sign[1];
                tmp++;
            }
            else if (remainder==9)
            {
                tmp[0] = sign[0];//X
                tmp[1] = sign[2];//I
                tmp+=2;
            }
            else {
                tmp[0] = sign[1];//V
                tmp+=1;
                remainder-=5;
                while(remainder){
                    tmp[0] = sign[0];//I
                    tmp++;
                    remainder--;
                }
            }
            
            

        }
    }
    tmp[0] = '\0';
    return res;

}

void main(void){
    int num = 58;
    printf("res is %s\n", intToRoman(num));
    
}
