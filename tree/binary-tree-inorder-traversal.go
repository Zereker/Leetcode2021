package main

import (
	"errors"
)

func Traversal(root *TreeNode, data *[]int) {
	if root == nil {
		return
	}

	Traversal(root.Left, data)
	*data = append(*data, root.Val)
	Traversal(root.Right, data)
}

func inorderTraversal2(root *TreeNode) []int {
	data := make([]int, 0)

	Traversal(root, &data)

	return data
}

var (
	ErrEmptyStack = errors.New("stack is empty")
)

type Stack struct {
	data []interface{}
}

func NewStack() *Stack {
	data := make([]interface{}, 0)

	return &Stack{data: data}
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s *Stack) Push(x interface{}) {
	s.data = append(s.data, x)
}

func (s *Stack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, ErrEmptyStack
	}

	index := len(s.data) - 1

	value := s.data[index]
	s.data = s.data[:index]

	return value, nil
}

func (s *Stack) Top() (interface{}, error) {
	if s.IsEmpty() {
		return nil, ErrEmptyStack
	}

	index := len(s.data) - 1

	value := s.data[index]

	return value, nil
}

func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)

	stack := NewStack()

	cur := root
	for cur != nil || !stack.IsEmpty() {
		if cur != nil {
			stack.Push(cur)
			root = cur.Left
		} else {
			pop, _ := stack.Pop()
			cur = pop.(*TreeNode)
			result = append(result, cur.Val)
			cur = cur.Right
		}
	}

	return result
}

func main() {
	inorderTraversal(&TreeNode{Val: 5, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 6}})
}
