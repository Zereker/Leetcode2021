package main

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	queue := NewQueue()
	if root != nil {
		queue.Enqueue(root)
	}

	for !queue.IsEmpty() {
		size := queue.Size()

		var (
			nodePre *Node
		)

		for i := 0; i < size; i++ {
			dequeue, _ := queue.Dequeue()
			node := dequeue.(*Node)

			if i == 0 { // 取出一层的头节点
				nodePre = node
			} else {
				nodePre.Next = node
				nodePre = node
			}

			if node.Left != nil {
				queue.Enqueue(node.Left)
			}

			if node.Right != nil {
				queue.Enqueue(node.Right)
			}
		}

		nodePre.Next = nil
	}

	return root
}
