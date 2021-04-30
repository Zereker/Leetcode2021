package main

/*
	https://leetcode-cn.com/problems/add-two-numbers/
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * @Description: 指针遍历
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := new(ListNode)
	curr := result
	addNum := 0
	for l1 != nil || l2 != nil || addNum > 0 {
		if l1 != nil {
			addNum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			addNum += l2.Val
			l2 = l2.Next
		}
		curr.Next = &ListNode{Val: addNum % 10}
		addNum /= 10
		curr = curr.Next
	}
	return result.Next
}
