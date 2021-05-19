package main

type NNode struct {
	Val      int
	Children []*NNode
}

func maxDepthNNode(root *NNode) int {
	if root == nil {
		return 0
	}

	var depth int
	for _, child := range root.Children {
		depth = max(depth, maxDepthNNode(child))
	}

	return depth + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
