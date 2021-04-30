package main

/*
	https://leetcode-cn.com/problems/add-two-numbers/
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		dummy = new(ListNode)
		curr  = dummy

		carry int
	)

	for l1 != nil || l2 != nil || carry > 0 {
		curr.Next = new(ListNode)

		curr = curr.Next
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}

		curr.Val = carry % 10
		carry /= 10
	}

	return dummy.Next
}

func main() {
	addTwoNumbers(
		&ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}},
		&ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}},
	)
}
