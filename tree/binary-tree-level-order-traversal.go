package main

import "errors"

var (
	ErrEmptyQueue = errors.New("queue is empty")
)

type Queue struct {
	data []interface{}
}

func NewQueue() *Queue {
	data := make([]interface{}, 0)

	return &Queue{data: data}
}

func (q *Queue) Enqueue(value interface{}) {
	q.data = append(q.data, value)
}

func (q *Queue) Dequeue() (interface{}, error) {
	if !q.IsEmpty() {
		value := q.data[0]
		q.data = q.data[1:]
		return value, nil
	}

	return nil, ErrEmptyQueue
}

func (q *Queue) Size() int {
	return len(q.data)
}

func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

func levelOrder(root *TreeNode) [][]int {
	result := make([][]int, 0)

	q := NewQueue()
	if root != nil {
		q.Enqueue(root)
	}

	for !q.IsEmpty() {
		size := q.Size()
		value := make([]int, 0, size)

		for i := 0; i < size; i++ {
			dequeue, _ := q.Dequeue()
			node := dequeue.(*TreeNode)

			value = append(value, node.Val)

			if node.Left != nil {
				q.Enqueue(node.Left)
			}

			if node.Right != nil {
				q.Enqueue(node.Right)
			}
		}

		result = append(result, value)
	}

	return result
}
