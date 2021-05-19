package main

func depth(node *TreeNode) int {
	if node == nil {
		return 0
	}

	left := depth(node.Left)
	right := depth(node.Right)

	if left > right {
		return left + 1
	}

	return right + 1
}

func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return depth(root)
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := NewQueue()
	queue.Enqueue(root)

	var depth int
	for !queue.IsEmpty() {
		depth++

		size := queue.Size()
		for i := 0; i < size; i++ {
			dequeue, _ := queue.Dequeue()
			node := dequeue.(*TreeNode)

			if node.Left != nil {
				queue.Enqueue(node.Left)
			}

			if node.Right != nil {
				queue.Enqueue(node.Right)
			}
		}
	}

	return depth
}
