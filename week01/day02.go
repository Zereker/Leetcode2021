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
		curr   = new(ListNode) // 当前节点, 第一次运行的时候是头节点
		result = curr          // 结果
		carry  int             // 是否进位
	)

	for l1 != nil || l2 != nil || carry > 0 {
		if l1 != nil {
			carry += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}

		// 下一次操作只有有意义再进行
		curr.Val = carry % 10
		if l1 != nil || l2 != nil || carry >= 10 {
			curr.Next = new(ListNode)
		}

		// 准备下一次循环
		carry = carry / 10
		curr = curr.Next
	}

	return result
}

func main() {
	addTwoNumbers(
		&ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}},
		&ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}},
	)
}
