#include <stdio.h>
/* 与官方不相伯仲，并略胜一筹 */
int maxProfit(int* prices, int pricesSize) // C没有办法知道一个数列的场长度，哈哈，老子笑了
{
    if (pricesSize<2) return 0;
    int full = 0;
    int cost = 0;
    int ret = 0;
    int i;
    for (i=0; i<pricesSize-1; i++)
    {
        if (prices[i]<prices[i+1])
        {
            if (!full)
            {
                cost = prices[i];
                // printf("choose %d as cost: %d\n", i, cost);
                full = !full;
            }
        }
        else
        {
            if (full)
            {
                ret += prices[i] - cost;
                // printf("choose %d to sell", i);
                // printf("ret is %d\n", ret);
                full = !full;

            }
        } 
    }
    if (full)
    {
        ret += prices[i] - cost;
        // printf("at last choose %d as sell price: %d\n", i-1, prices[i-1]);
    }

    return ret;
    
}
#define len 2
void main(void)
{
    int a[len] = {2,4};
    int length = len;
    int res;
    res = maxProfit(a, length);
    printf("res is %d", res);

}