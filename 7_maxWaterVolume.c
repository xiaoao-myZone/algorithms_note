#include <stdio.h>
#include <string.h>
#include <stdbool.h>

/*
Pre:
    1. 题目中将这个最大面积具象成为灌水问题,有一点迷惑性
    2. 用最不烧脑的方法来想,就是排列组合这n的数,算出所有的面积取最大值
    3. 但是这些排列组合中一定有一些是不用去尝试的
    4. 如果从数列的两边向中间试探,那么最大高度不大于已经尝试过的组合就不必去计算,直接跳过
    5. 但是从中间向两边尝试,就没有这个效果
    6. 试探需要限定一个最大宽度,在这个宽度内,需要从左至右(反之也行)做一次扫描,然后再缩小这个宽度,重复这个操作
    7. 如果按照6的想法做下去,会去每种可能,至少要做一次取短边的比较
    8. 应该更灵活一点,一次只移动一边并且是短边移动
    9. 证明题: 这个数组中存在最靠边的最大值(最大值可以有多个,加上前一个条件,必有两个),那水一定会接触这两个最大值
    10. 9比较困难,我们可以简单证明:当唯一最大值出现在最边上时,它一定是其中一个边界
    11. 试证明,当存在唯一最大值时,水一定接触它

            ①假设m处的值最大,且存在a,b两处包围的面积最大(a<b), a,b均不等于m
            max = f(a)(b-a)或f(b)(b-a), 不妨设f(a)<f(b)
            若将m与b纳入考虑
            则有area = f(b)|b-m|
                1. 当m<b时:
                    必须有b-m>b-a,即m<a时,假设才不成立,反证假设错误
                2. 当m>b时:
                    必须有m-b>b-a,即m>2b-a时, 假设才不成立,反证假设错误
            ②当b>m>a,m在a,b 之间必被水接触
            ③或b<m<2b-a时,此时将m与a纳入考虑,必然f(a)(m-a)>max,假设才不成立,反证假设错误

            所以综合上述,存在唯一最大值时,要么该最大值作为边界,要么在a,b之间
    12. 再尝试证明
    13. 貌似这个证明对解题没什么用




Post:
    1. 双指针将原本的O(n2)变为O(n)
    2. 当两个边界相等时,可以同时推进

*/

int max(int a, int b){
    if (a>b) return a;
    else return b;
}

int min(int a, int b){
    if (a>b) return b;
    else return a;
}

int maxArea_1(int* height, int heightSize){
    int min_margin, max_volume;
    min_margin = max_volume = 0;
    int left, right;
    int range = heightSize-1;
    int tmp;
    while (range>0){
        for(left=0, right=range; right<heightSize; left++,right++){
            //printf("left is %d, right is %d\n", left, right);
            tmp = min(height[left], height[right]);
            if (tmp>min_margin){
                max_volume = max(tmp*(right-left), max_volume);
                min_margin = tmp;
                //printf("**update max_volume is %d, min_margin is %d\n", max_volume, min_margin);
            }

        }
        range--;
    }
    printf("last num is %d", height[right]);
    return max_volume;

}

/*
执行用时：96 ms, 在所有 C 提交中击败了31.93% 的用户
内存消耗：11.5 MB, 在所有 C 提交中击败了15.91% 的用户  有点进步
*/
int maxArea_2(int* height, int heightSize){
    int max_volume=0;
    int left, right, tmp;
    left = tmp = 0; right = heightSize-1;
    bool side = false; //left is false, right is true
    int left_tmp = 0;
    int right_tmp = 0;
    printf("right is %d\n", right);
    while (left!=right){

        if (side){
            if (height[right]>tmp){}
            else {
                right--;
                continue;
            }
        }
        else {
            if (height[left]>tmp){}
            else {
                left++;
                continue;
            }
        }
        printf("right is %d\n", right);
        if (height[left]>height[right]){
            printf("update left: %d, right: %d\n", left, right);
            max_volume = max(height[right]*(right-left), max_volume);
            printf("update left_h: %d, right_h: %d\n", height[left], height[right]);
            tmp = height[right];
            side = true;
            right--;
        }
        else if (height[left]<height[right])
        {
            printf("update left: %d, right: %d\n", left, right);
            max_volume = max(height[left]*(right-left), max_volume);
            printf("update left_h: %d, right_h: %d\n", height[left], height[right]);
            tmp = height[left];
            side = false;
            left++;
        }
        else {//height[left]==height[right]
            printf("update left: %d, right: %d\n", left, right);
            max_volume = max(height[left]*(right-left), max_volume);
            printf("update left_h: %d, right_h: %d\n", height[left], height[right]);
            tmp = height[left];

            left_tmp = left;
            left_tmp++;
            while (left_tmp!=right) {
                if (height[left_tmp]>height[left]){
                    break;
                }
                else {
                    left_tmp++;
                }
            }

            right_tmp = right;
            right_tmp--;
            while (right_tmp!=left) {
                if (height[right_tmp]>height[right]){
                    break;
                }
                else {
                    right_tmp--;
                }
            }

            if (right-right_tmp<left_tmp-left){ //右边先找到一个较大的值
                right = right_tmp;
                side = true;//理论上说这个有没有无所谓
            }
            else {//如果相等,让左边先走也无妨,因为下一轮算完面积后,会右边会先行,但是如果标记这个步骤,就不用在重复比较了
                left = left_tmp;
                side = false; //理论上说这个有没有无所谓
            }
        }
        
    }

    return max_volume;
}

/*
执行用时：84 ms, 在所有 C 提交中击败了36.69% 的用户
内存消耗：11.4 MB, 在所有 C 提交中击败了24.62% 的用户
*/
int maxArea(int* height, int heightSize){
    int max_volume=0;
    int left, right, tmp;
    left = tmp = 0; right = heightSize-1;
    bool side = false; //left is false, right is true
    printf("right is %d\n", right);
    while (left<right){

        if (side){
            if (height[right]>tmp){}
            else {
                right--;
                continue;
            }
        }
        else {
            if (height[left]>tmp){}
            else {
                left++;
                continue;
            }
        }
        printf("right is %d\n", right);
        if (height[left]>height[right]){
            printf("update left: %d, right: %d\n", left, right);
            max_volume = max(height[right]*(right-left), max_volume);
            printf("update left_h: %d, right_h: %d\n", height[left], height[right]);
            tmp = height[right];
            side = true;
            right--;
        }
        else if (height[left]<height[right])
        {
            printf("update left: %d, right: %d\n", left, right);
            max_volume = max(height[left]*(right-left), max_volume);
            printf("update left_h: %d, right_h: %d\n", height[left], height[right]);
            tmp = height[left];
            side = false;
            left++;
        }
        else {//height[left]==height[right]
            printf("update left: %d, right: %d\n", left, right);
            max_volume = max(height[left]*(right-left), max_volume);
            printf("update left_h: %d, right_h: %d\n", height[left], height[right]);
            tmp = height[left];
            //如果两个相等高度的范围内存在更大的面积,那么一定不以这两个地方为边界
            left++;
            right--;

 
        }
        
    }

    return max_volume;
}

void main(void)
{
    int nums[] = {1,1};
    printf("result is %d\n", maxArea(nums, 2));
    // 20000 --> 197991309
}