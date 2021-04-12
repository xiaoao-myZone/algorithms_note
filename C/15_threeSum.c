#include <stdio.h>
#include <malloc.h>
#include <stdbool.h>

/*

*/

/*
pre:
    1. 暴力的做法是从n个数中抽3个进行,这样时间复杂度是O(n3)
    





    3. 如果先排序nlog(n),可以变成O(n2)
    4. 注意: 列表里存在重复的数,但是要求结果不可以存在重复的结果
    5. 排序后, 结果是全零,或者有正有负
    6. 排序后, 可以用二分法,双指针

post:
    1.

*/
int * twoSum(int * nums, int size){//默认和为0
    int tmp;
    for(int i=0; i<size; i++){
        tmp = 0-nums[i];
        for(int j=i+1; j<size; j++){
            if (tmp==nums[j]){
                int * ret = (int *)malloc(sizeof(int)*2);
                ret[0] = i;
                ret[1] = j;
                return ret;
            }
        }
    }
    return NULL;
}

int** old_threeSum(int* nums, int numsSize, int* returnSize, int** returnColumnSizes){
    int ** ret;
    ret = NULL;
    int fisrt = 0;
    int i=0;
    for(i=0; i<numsSize; i++){//使用两数之和还需要跳过选中的列表, 最好是在twoSum中分成两段
        

    }

    ret = (int **)malloc(sizeof(int *)*(*returnSize));
    returnColumnSizes = (int **)malloc(sizeof(int *)*(*returnSize));
    for(i=0;i<*returnSize;i++){
        ret[0] = (int *)malloc(sizeof(int) * 3);
        returnColumnSizes[0] = NULL;
    }

    return ret;

}

void bubbleSort(int * nums, int size){
    int tmp;
    for(int j=size-1; j>0; j--){
        for(int i=0; i<j; i++){
            if (nums[i]>nums[i+1]){
                tmp = nums[i];
                nums[i] = nums[i+1];
                nums[i+1] = tmp;
            }
        }
    }
}

int** threeSum(int* nums, int numsSize, int* returnSize, int** returnColumnSizes){
    bubbleSort(nums, numsSize); // ascent
    int ** ret;
    //ret = (int**)malloc(sizeof(int *)*1);
    int i,j;
    i = 0;
    *returnSize = 0;
    j=numsSize;
    int tmp;
    int third;
    bool isCatch = false;
    while (i!=j){
        isCatch = false;
        third = -(nums[i] + nums[j]);
        if (third>0){
            for(tmp=j-1; nums[tmp]>0; tmp--){
                if (nums[tmp]==third){
                    isCatch = true;
                    (*returnSize)++;
                    ret = (int **)realloc(ret, sizeof(int *)*(*returnSize));
                    //ret[(*returnSize)-1] = {nums[i], nums[tmp], nums[j]};
                    ret[(*returnSize)-1] = (int *)malloc(sizeof(int)*3);
                    ret[(*returnSize)-1][0] = nums[i];
                    ret[(*returnSize)-1][1] = nums[tmp];
                    ret[(*returnSize)-1][2] = nums[j];
                    break;
                }
            }
            if (!isCatch){
                
            }
        }
        else {
            for(tmp=i+1; nums[tmp]<=0; tmp++){
                if (nums[tmp]==third){
                    isCatch = true;
                    (*returnSize)++;
                    ret = (int **)realloc(ret, sizeof(int *)*(*returnSize));
                    //ret[(*returnSize)-1] = {nums[i], nums[tmp], nums[j]};
                    ret[(*returnSize)-1] = (int *)malloc(sizeof(int)*3);
                    ret[(*returnSize)-1][0] = nums[i];
                    ret[(*returnSize)-1][1] = nums[tmp];
                    ret[(*returnSize)-1][2] = nums[j];
                    break;
                }
            }
        }

    }
}


void main(void){
    int nums[] = {-1,0,1,2,-1,-4};
    int * returnSize;
    int ** returnColumnSizes;
    bubbleSort(nums, 6);
    for(int i=0;i<6;i++){
        printf("%d\n", nums[i]);
    }

    
}
