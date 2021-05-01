package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// 思路：有限计算当前节点的值并初始化，后面循环去计算，注意：当l1\l2都为空的时候，需要判断上一个节点的和的十位数是否>0，如满足需要将这个节点也添加进去
func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	node := new(ListNode) // 当前节点
	result := node // reslut
	carry := 0 // 进位

	for l1 != nil || l2!=nil || carry >0 {
		if l1!= nil {
			carry += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			carry += l2.Val
			l2 = l2.Next
		}

		node.Val=carry%10
		carry = carry/10

		if l1 != nil || l2 != nil || carry > 0 {
			node.Next = new(ListNode)
		}

		node = node.Next
	}
	return result
}
