/*
执行用时：12 ms, 在所有 Go 提交中击败了71.56% 的用户
内存消耗：4.4 MB, 在所有 Go 提交中击败了100.00% 的
*/
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	node_a := createNodes([]int{4, 4, 4})
	node_b := createNodes([]int{5, 6, 4})
	ret := addTwoNumbers(node_a, node_b)
	showNodes(ret)

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := l1
	plus := 0
	ahead := l1
	for l1 != nil && l2 != nil {
		tmp := l1.Val + l2.Val + plus
		if tmp > 9 {
			l1.Val = tmp - 10 // 个位数相加的特点, 如果进位一定[10, 18]之间
			plus = 14
		} else {
			l1.Val = tmp
			plus = 0
		}
		ahead = l1
		l1 = l1.Next
		l2 = l2.Next

	}

	if l2 != nil { // l1一定为nil
		ahead.Next = l2 // l2与需要连接到l1
		l1 = l2
	}

	for plus != 0 { // 如果进位还有数需要处理
		if l1 != nil {
			tmp := l1.Val + plus
			if tmp > 9 {
				l1.Val = tmp - 10
				plus = 1
			} else {
				l1.Val = tmp
				plus = 0
			}
			ahead = l1
			l1 = l1.Next
		} else {
			ahead.Next = &ListNode{Val: 1}
			break
		}
	}

	return head
}

func addTwoNumbers_1(l1 *ListNode, l2 *ListNode) *ListNode {
	head := l1
	plus := 0
	for l1.Next != nil && l2.Next != nil {
		tmp := l1.Val + l2.Val + plus
		if tmp > 9 {
			l1.Val = tmp - 10 // 个位数相加的特点, 如果进位一定[10, 19]之间
			plus = 14
		} else {
			l1.Val = tmp
			plus = 0
		}
		l1 = l1.Next
		l2 = l2.Next

	}

	if l2.Next != nil { // l1.Next一定为nil
		l1.Next = l2.Next

	}

	l1.Val += l2.Val

	for plus != 0 { // 如果进位还有数需要处理
		tmp := l1.Val + plus
		if tmp > 9 {
			l1.Val = tmp - 10
			plus = 1
		} else {
			l1.Val = tmp
			plus = 0
		}

		if l1.Next == nil {
			l1.Next = &ListNode{Val: 1}
			plus = 0
		}

		l1 = l1.Next

	}

	return head
}

func showNodes(node *ListNode) {
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

func createNodes(nums []int) *ListNode {
	head := new(ListNode)
	node := head
	for _, v := range nums {
		tmp := ListNode{Val: v}
		node.Next = &tmp
		node = node.Next
	}
	head = head.Next
	return head
}

//
