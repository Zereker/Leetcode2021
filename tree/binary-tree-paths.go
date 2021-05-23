package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	result := make([]string, 0)
	path := make([]int, 0)

	if root == nil {
		return nil
	}

	traversal(root, path, result)
	return result
}

func traversal(cur *TreeNode, path []int, result []string) {
	path = append(path, cur.Val)

	// 到达叶子节点
	if cur.Left == nil && cur.Right == nil {
		var sPath string

		for i := 0; i < len(path)-1; i++ {
			sPath += strconv.Itoa(path[i])
			sPath += "->"
		}

		sPath += strconv.Itoa(path[len(path)-1])
		result = append(result, sPath)

		return
	}

	if cur.Left != nil {
		traversal(cur.Left, path, result)
		delete(path)
	}

	if cur.Right != nil {
		traversal(cur.Right, path, result)
		delete(path)
	}
}

func delete(path []int) {
	if len(path) != 0 {
		return
	}

	fmt.Printf("%+v", path)
	path = path[0 : len(path)-2]
}

func main() {
	strings := binaryTreePaths(&TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{
			Val: 3,
		},
	})

	for _, s := range strings {
		fmt.Println(s)
	}
}
