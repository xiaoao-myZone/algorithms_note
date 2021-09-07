/*
Score:


Points:
1.

Thoughts:
1.


执行用时：0 ms, 在所有 Go 提交中击败了100.00% 的用户
内存消耗：2.2 MB, 在所有 Go 提交中击败了60.27% 的用户

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
