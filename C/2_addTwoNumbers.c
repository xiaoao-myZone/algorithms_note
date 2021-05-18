#include <stdio.h>
#include <malloc.h>
/* 惨败，时间与空间垫底 */
/*
几个月后再次尝试
执行用时：16 ms, 在所有 C 提交中击败了80.36% 的用户
内存消耗：7 MB, 在所有 C 提交中击败了94.77% 的用户

又过了几个月后继续尝试
执行用时：8 ms, 在所有 C 提交中击败了99.27% 的用户
内存消耗：7.2 MB, 在所有 C 提交中击败了87.66% 的用户
*/
struct ListNode {
     int val;
     struct ListNode *next;
};

struct ListNode* Old_addTwoNumbers(struct ListNode* l1, struct ListNode* l2)
{
    struct ListNode* ret = (struct ListNode *)malloc(sizeof(struct ListNode));
    ret->next = NULL;
    int tmp = 0;
    printf("first tmp is %d\n", tmp);
    struct ListNode* head = ret;
    while (l1 && l2)
    {
        printf("nums are %d and %d\n", l1->val, l2->val);
        ret->next = (struct ListNode *)malloc(sizeof(struct ListNode));
        ret = ret->next;
        ret->next = NULL;
        tmp += l1->val + l2->val;
        ret->val = tmp % 10;
        tmp = tmp>=10;
        l1 = l1->next;
        l2 = l2->next;
        printf("tmp is %d\n", tmp);
    }

    while (l1)
    {
        ret->next = (struct ListNode *)malloc(sizeof(struct ListNode));
        ret = ret->next;
        ret->next = NULL;
        tmp += l1->val;
        ret->val = tmp % 10;
        tmp = tmp>=10;
        l1 = l1->next;
    }
    while (l2)
    {
        ret->next = (struct ListNode *)malloc(sizeof(struct ListNode));
        ret = ret->next;
        ret->next = NULL;
        tmp += l2->val;
        ret->val = tmp % 10;
        tmp = tmp>=10;
        l2 = l2->next;
    }
    if (tmp)
    {
        ret->next = (struct ListNode *)malloc(sizeof(struct ListNode));
        ret = ret->next;
        ret->next = NULL;
        ret->val = tmp;
    }
    return head->next;

}

// new
// no limit to save l1 and l2

struct ListNode* addTwoNumbers(struct ListNode* l1, struct ListNode* l2){
    struct ListNode * head;
    head = l1;
    unsigned i = 0;
    while (l1->next && l2->next){
        l1->val += l2->val + i;
        //i = l1->val / 10;
        // l1->val %= 10;
        if (l1->val>9) {
            i=1;
            l1->val -= 10;
        }
        else {
            i=0;
        }
        l1 = l1->next;
        l2 = l2->next; //如果这样,当两个链表长度相同时,将获取不到最后的句柄
    }

    l1->val += l2->val + i;
    //i = l1->val / 10;
    // l1->val %= 10;
    if (l1->val>9) {
        i=1;
        l1->val -= 10;
    }
    else {
        i=0;
    }

    if(l2->next){//l1指向最后一个数

        l1->next = l2->next; //嫁接  
    }
    while (i){
        
        if(l1->next){
            l1 = l1->next;//上一位已经操作完成
            l1->val += i;
            //i = l1->val / 10;
            // l1->val %= 10;
            if (l1->val>9) {
                i=1;
                l1->val -= 10;
            }
            else {
                i=0;
            }
        }
        else{
            l1->next = (struct ListNode *)malloc(sizeof(struct ListNode));
            l1 = l1->next;
            l1->next = NULL;
            l1->val = i; //i不可能超过1
            break;
            
        }
    } 
    return head;   
}

void print_chain(struct ListNode *node){
    while(node){
        printf("%d", node->val);
        node = node->next;
    }
    printf("\n");
}

void main(void)
{
   struct ListNode *l1 = (struct ListNode *)malloc(sizeof(struct ListNode));
   l1->val = 1;
   struct ListNode * head1 = l1;
   struct ListNode *l2 = (struct ListNode *)malloc(sizeof(struct ListNode));
   l2->val = 2;
   struct ListNode * head2 = l2;
   for (int i=5; i<8; i++)
   {
       l1->next = (struct ListNode *)malloc(sizeof(struct ListNode));
       l1 = l1->next;
       l1->val = i;
       l1->next = NULL;
       l2->next = (struct ListNode *)malloc(sizeof(struct ListNode));
       l2 = l2->next;
       l2->val = i;
       l2->next = NULL;
   }
   struct ListNode *res;
   //head1 :1567 head2:2567

   
   print_chain(head1);
   print_chain(head2);
   res = addTwoNumbers(head1, head2);
   print_chain(res);


}

/*
Conclusion
1. 移花接木,赏心悦目
2. 当长度不明确时,链表最好以p->next为循环条件
*/
