#include <stdio.h>
#include <string.h>
#include <malloc.h>
#include <stdbool.h>

void main(void){
    /*1. 求余与地板除*/
    // int i = 14;
    // i %= 10;
    // i += 1+1;
    // printf("%d\n", i);
    // printf("%d\n", 10/2);
    // printf("%d\n", 10%2);

    /*2. * 与++的优先级*/
    //char s[6] = "xiaoao";//超范围直接返回0?windows上是现实一个随机数
    // char *s = "xiaoao";//这种定义方式才能使用*s++
    // printf("aa%c\naa\n", s[5]);
    // while (*s){
    //     printf("%c", *s++); //$$$$$$$$先取值,再++$$$$$$$$$$$
    // }
    // printf("\n");

    /*3. strlen 与sizeof*/
    // char *s = "ookkllll";
    // printf("sizeof: %ld\n", sizeof(s)); //$$$$$$$ s是一个指针$$$$$$$$$$$$
    // printf("strlen: %ld\n", strlen(s));
    /*3. s是一个指针*/

    /*4. *与++/--的优先级 */
    // int i = 0;
    // int * p = &i;
    // printf("i is %d\n", (*p)++); //*p++
    // printf("i is %d\n", i);
    // printf("i is %d\n", *p);
    /* *运算符的优先级大于++/-- */
    /* TODO *p++是先取值,再返回给%d, 然后对p++ 类似 a = *p++ */

    /*5. *与==的优先级 */
    // int i = 0;
    // int * p = &i;
    // printf("answer is %d\n", *p == 0);
    /* *的优先级大于 == */

    /*6. 字符串的分片 */
    // char * s = "abcdef";
    // int len = 2;
    // printf("strlen: %ld\n", strlen(s));
    // char * des = (char *)malloc(sizeof(char)*(len));

    // strncpy(des, s+3, len);
    // //strcat(des, s);
    // printf("%s|\n", des);

    /*7. 或与非*/
    // int i=0, j=2;
    // printf("ret: %d\n", i>-1&&j<2);
    // printf("ret: %d\n", i>-1||j<2);

    /*8. 修改初始字符*/
    //char * s = "asdasf";
    // char s[20] = "asdasf\0";
    // s[3]='P';
    // printf("%s|\n", s);

    /*9. 修改头指针*/
    // char s[20] = "asdasf\0";
    // //char * s = "asdasf";
    // char * p;
    // p = s;
    // p++;
    // printf("%s|\n", p);

    /*10. 字符串赋值*/
    // char *s;
    // s = "asdasfa";
    // int len = strlen(s);
    // printf("%s|\n", s);
    // printf("ret:%d\n", s[len]=='\0');

    /*11. */
    // bool a=true;
    // printf("%ld\n", sizeof(a));

    /*12. 求余*/
    // int num = 5;
    // printf("%d\n", num%10/(10/10));

    /*13. realloc*/
    // int ** ret;
    // //ret = (int **)malloc(sizeof(int *)*1);
    // ret = (int **)realloc(ret, sizeof(int *)*2);
    // int a[] = {1,2,3};
    // ret[0] = a;
    // ret = (int **)realloc(ret, sizeof(int *)*3);
    // ret[2] = a;
    // printf("%d, %d\n", ret[2][1], ret[2][0]);

    /*14. point and array name */
    int a[]={1,2,3,4,5};
    int *p = a;
    p++;
    // a++; syntax error
    printf("%d\n", p[1]);
    
    


}