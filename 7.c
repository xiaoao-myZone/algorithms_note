#include <stdio.h>
#include <string.h>
#include <stdbool.h>
#include <limits.h>

int myAtoi(char * s){
    int start=-1;
    int cnt = 0;
    char tmp;
    int i;
    long ret=0;
    bool sign = false;
    char * overflow = "2147483648";
    for (i=0;i<strlen(s);i++)
    {
        tmp = s[i];


        if (47<tmp && tmp<58)
        {
            if (start<0) start = i;
            cnt++;
        }
        else
        {
            if (start<0)
            {
                if (tmp==' ') continue;
                else
                {
                    if (!sign && (tmp=='-' || tmp=='+'))
                    {
                        sign = true;
                        continue;
                    }
                    else return 0;
                }
            }
            else
            {
                break;
            }
        }
            
        
    }
    // printf("start: %d\n", start);
    // printf("cnt: %d\n", cnt);
    // if (cnt > 10) {
    //     if (start>0 && s[start-1]=='-') return -2147483648;
    //     else return 2147483647;
    // }
    // else if (cnt == 10)
    // {
    //     int j=0;
    //     char tmp_1;
    //     for (int end=start+cnt,i=start;i<end; i++,j++)
    //     {
    //         tmp = s[i];
    //         tmp_1 = overflow[j];
    //         if (tmp>tmp_1) {
    //             if (start>0 && s[start-1]=='-') return -2147483648;
    //             else return 2147483647;
    //         }
    //         else if (tmp==tmp_1) continue;
    //         else break;
    //     }
    // }
    if (false) return 0;
    else
    {
        for (int end=start+cnt,i=start; i<end; i++)
        {
            // printf("%d\n" ,s[i]);
            // printf("ret: %lld\n", ret);
            ret = 10 * ret + s[i] - '0';
            
        }
        if (start>0 && s[start-1]=='-') 
        {
            ret*=-1;
            if ((int)ret != ret) 
            {
                return INT_MIN;

            }
        }
        else
        {
            if ((int)ret != ret) 
            {
                return INT_MAX;

            }
        }
        
        
    }
    return ret;

}

void main(void)
{
    char * s = "20000000000000000000";
    // "-91283472332" --> -2147483648 
    // "21474836460" --> 2147483647
    // "  0000000000012345678" --> 12345678
    // 20000000000000000000"
    printf("\n%d", myAtoi(s));
}