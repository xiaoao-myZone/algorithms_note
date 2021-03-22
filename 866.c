// lemon water
#include <stdio.h>
#include <stdbool.h>
// count
bool lemonadeChange(int* bills, int billsSize)
{
    int b_5, b_10;
    b_5=b_10=0;
    for (int i=0; i<billsSize; i++)
    {
        if (bills[i]==5)
        {
            b_5 += 1;
        }
        else 
        {
            if (bills[i]==10)
            {
                if (b_5>0)
                {
                        b_5 -= 1;
                        b_10 += 1;
                }
                else return false;
            }
            else 
            {
                if (b_10>0 && b_5>0)
                {
                    b_10 -= 1;
                    b_5 -= 1;
                }
                else 
                {
                    if (b_5>2)
                    {
                        b_5 -= 3;
                    } 
                    else return false;
                }
            }
        }
    }
    return true;

}


bool lemonadeChange(int* bills, int billsSize)
{
    int tmp = 0;
    for (int i=0; i<billsSize; i++)
    {
        if (bills[i]==5) continue;
        else if (bills[i]==10) {
            tmp = i;

        }
        else {

        }
    }
}
void main(void)
{
    int array[5] = {5,5,5,10,20};
    printf("%d" , lemonadeChange(array, 5));

}
