#include <stdio.h>
/* 82 28 */
double findMedianSortedArrays(int* nums1, int nums1Size, int* nums2, int nums2Size){
    int sum = nums1Size+nums2Size;
    int nums[sum];
    nums1Size--;
    nums2Size--;
    int i;
    for (i=0; nums1Size>=0 && nums2Size>=0; i++)
    {

        if (nums1[nums1Size] > nums2[nums2Size])
        {
            nums[i] = nums1[nums1Size];
            nums1Size--;
        }
        else
        {
            nums[i] = nums2[nums2Size];
            nums2Size--;
                
        }
    }
            
    while (nums1Size>=0)
    {
        nums[i] = nums1[nums1Size];
        nums1Size--;
        i++;
    }
    while (nums2Size>=0)
    {
            nums[i] = nums2[nums2Size];
            nums2Size--;
            i++;
    }
    // for(i=0; i<sum;i++)
    // {
    //     printf("%d\n", nums[i]);
    // }
    int mid = sum / 2;
    if (sum%2)
    {
        return nums[mid];
    }
    else
    {
        return (nums[mid]+nums[mid-1])/2.0;
    }
    

}
void main(void)
{
    int num1[3] = {1,3,5};
    int num2[5] = {0,2,4,6,8};
    float res = findMedianSortedArrays(num1, 3, num2, 5);
    printf("%f", res);
}