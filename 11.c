#include <stdio.h>
/*
1. 这种问题难点在于不可能线性地去遍历从而得到结果，需要程序具有探测能力，就是既要立足当前，还要适当往前
2. 右侧应该从大到小，这样可以顺便过滤一些操作

执行用时：716 ms, 在所有 C 提交中击败了17.50% 的用户
内存消耗：7.2 MB, 在所有 C 提交中击败了5.09% 的用户
QVQ
*/
int maxArea(int* height, int heightSize){
    int tmp=0;
    int ret=0;
    int left,right;
    int tmp_left=0;
    int tmp_right = 0;
    //int cnt = 0;
    for (left=0;left<heightSize-1;left++){
        //printf("left index:%d -- height:%d -- tmp:%d\n", left, height[left], tmp_left);
        if (height[left]<=tmp_left) {
            //printf("left pass\n");
            continue;
            }
        tmp_left = height[left];
        for (right=heightSize-1,tmp_right=0;right>left;right--){
            //printf("right index:%d -- height:%d -- tmp:%d\n", right, height[right], tmp_right);
            if (height[right]>tmp_right){
                tmp_right = height[right];
                //cnt++;
                //printf("%d -- %d\n", left, right);
                tmp = (height[left]>height[right]?height[right]:height[left]) * (right-left);
                if (tmp>ret){
                    ret = tmp;
                }
            }
            else{
                //printf("right pass %d -- %d\n", height[right], right);
                continue;
            }
        }
    }
    //printf("count:%d\n", cnt);
    return ret;
}

void main(void)
{
    int h[] = {1,8,6,2,5,4,8,3,7};
    printf("%d", maxArea(h,9));

}