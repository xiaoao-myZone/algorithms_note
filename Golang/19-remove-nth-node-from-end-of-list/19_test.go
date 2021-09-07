package leetcode

import (
	"fmt"
	"testing"
)

type Input struct {
	nodes []int
	index int
}
type TestCase struct {
	input  Input
	output []int
}

func creat_node(nodes []int) *ListNode {
	ret := &ListNode{0, nil}
	tmp := ret
	for _, val := range nodes {
		tmp.Next = &ListNode{val, nil}
		tmp = tmp.Next
	}
	return ret.Next
}

func check_node(head *ListNode, nodes []int) bool {
	for _, val := range nodes {
		fmt.Println(head.Val, val)
		if head.Val == val {
			head = head.Next
			continue
		} else {

			return false
		}

	}
	if head == nil {
		return true
	} else {
		return false
	}

}
func Test_0_template(t *testing.T) {
	test_cases := []TestCase{
		{Input{[]int{1, 2, 3, 4, 5}, 2}, []int{1, 2, 3, 5}},
		{Input{[]int{1}, 1}, []int{}},
		{Input{[]int{1, 2}, 2}, []int{2}},
	}
	for _, test_case := range test_cases {
		head := creat_node(test_case.input.nodes)
		if res := removeNthFromEnd(head, test_case.input.index); check_node(res, test_case.output) {
			t.Log("通过")
		} else {
			t.Error("未通过", res, "/", test_case.output)
		}
	}
}
