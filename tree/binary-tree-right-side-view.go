package main

func rightSideView(root *TreeNode) []int {
	q := NewQueue()

	if root != nil {
		q.Enqueue(root)
	}

	result := make([]int, 0)
	for !q.IsEmpty() {
		size := q.Size()

		for i := 0; i < size; i++ {
			dequeue, _ := q.Dequeue()
			node := dequeue.(*TreeNode)

			if i == size-1 {
				result = append(result, node.Val)
			}

			if node.Left != nil {
				q.Enqueue(node.Left)
			}

			if node.Right != nil {
				q.Enqueue(node.Right)
			}
		}
	}

	return result
}
