#include <stdio.h>
#include <malloc.h>
// 在一个数组中找到两个数的和等于目标值
/* 垃圾leecode 拷贝参考答案提交都超时 */
int* twoSum(int* nums, int numsSize, int target, int* returnSize){
    //int * ret = (int *)malloc(sizeof(int));
    static int ret[2];
    int another;
    for (int i=0; i<numsSize-1; i++)
    {
        another = target - nums[i];
        ret[0] = i;
        for (int j=i+1; j<numsSize; j++)
        {
            if (nums[j]==another)
            {
                ret[1] = j;
                *returnSize = 2;
                return ret;
            }
        }
    }
    *returnSize = 0;
    return NULL;

}


#define len 5
void main(void)
{
    int nums[len] = {1,7,3,4,5};
    int target = 7;
    int * size;
    int * res;
    res = twoSum(nums, len, target, size);
    printf("index is %d and %d\n", res[0], res[1]);
}

// Conclusion: 两两结合还是,从结果找答案?