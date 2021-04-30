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

	l1Val, l2Val := 0, 0

	if l1 != nil {
		l1Val = l1.Val
		l1 = l1.Next
	}

	if l2 != nil {
		l2Val = l2.Val
		l2 = l2.Next
	}

	val := (l1Val + l2Val) % 10
	item := (l1Val + l2Val) / 10

	res := ListNode{Val: val, Next: nil}
	head := &res
	for {
		if l1 == nil && l2 == nil {
			if item == 0 {
				break
			} else {
				node := ListNode{Val: item, Next: nil}
				head.Next = &node
				break
			}

		}

		l1Val, l2Val = 0, 0

		if l1 != nil {
			l1Val = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			l2Val = l2.Val
			l2 = l2.Next
		}

		val := (l1Val + l2Val + item) % 10
		item = (l1Val + l2Val + item) / 10

		node := ListNode{Val: val, Next: nil}
		head.Next = &node
		head = head.Next
	}

	return &res
}
