package main

func zigzagLevelOrder(root *TreeNode) [][]int {
	var (
		level  int
		result = make([][]int, 0)
		queue  = NewQueue()
	)

	if root != nil {
		level = 1
		queue.Enqueue(root)
	}

	for !queue.IsEmpty() {
		size := queue.Size()
		values := make([]int, 0, size)

		for i := 0; i < size; i++ {
			dequeue, _ := queue.Dequeue()
			node := dequeue.(*TreeNode)

			values = append(values, node.Val)

			if node.Left != nil {
				queue.Enqueue(node.Left)
			}

			if node.Right != nil {
				queue.Enqueue(node.Right)
			}
		}

		if level%2 == 0 { // 奇数顺序遍历, 偶数逆序遍历
			tmp := make([]int, 0, len(values))

			for i := len(values) - 1; i >= 0; i-- {
				tmp = append(tmp, values[i])
			}

			values = tmp
		}

		level++
		result = append(result, values)
	}

	return result
}
