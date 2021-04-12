#include <stdio.h>
#include <stdbool.h>
#include <string.h>

bool isMatch(char * s, char * p){
    char tmp;
    int i=0;
    int dup = false;
    int p_len = strlen(p);
    int s_len = strlen(s);
    if (!p_len)
    {
        if (!s_len) return true;
        else  return false;
    }

    tmp = *p++;
    dup = (*p=='*');
    while (*s != '\0'){
        if (tmp!='.'){
            if (tmp!=*s){
                if (dup){
                    dup = false;
                    p++;
                    } 
                
                else{
                    return false;
                }
            }
            else{
               s++;
            }
        }
        else{
            s++;
        }
        // if (dup){
        //     dup = false;
        //     } 
        if (dup) continue;
        if (*p!='\0'){
            tmp = *p++;
            dup = (*p=='*');
        }
        else{
            tmp = '\0';
        }
    }
    return true; 
        

}

void main(void)
{
    char *s = "aab";
    char *p = "c*a*b";
    printf("%d", isMatch(s,p));

}