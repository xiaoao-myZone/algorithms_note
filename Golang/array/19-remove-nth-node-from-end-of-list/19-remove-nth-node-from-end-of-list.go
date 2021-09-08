/*
Score:
执行用时：0 ms, 在所有 Go 提交中击败了100.00% 的用户
内存消耗：2.2 MB, 在所有 Go 提交中击败了60.27% 的用户

Points:
1.

Thoughts:
1.


给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

进阶：你能尝试使用一趟扫描实现吗？



示例 1：

输入：head = [1,2,3,4,5], n = 2
输出：[1,2,3,5]


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。




*/
package leetcode

// import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 优化： 当n=1时， 直接不需要tmp
	tmp, current := head, head
	// 与current保持n个距离，当current到底时， tmp在目标节点的上一个节点
	for i := 0; i < n; i++ {
		current = current.Next
	}
	// 当长度等于n的时候就会出现current==nil，此时应该抽取第一个就可以了
	if current == nil {
		return tmp.Next
	}
	for current.Next != nil {
		tmp = tmp.Next
		current = current.Next
	}
	// 删除tmp.Next
	current = tmp.Next.Next
	tmp.Next = current
	return head
}
