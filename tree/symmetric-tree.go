package main

// 单层递归的逻辑就是处理左右节点都不为空，且数值相同的情况
func compare(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}

	if left == nil || right == nil {
		return false
	}

	if left.Val != right.Val {
		return false
	}

	return compare(left.Left, right.Right) && compare(left.Right, right.Left)
}

func isSymmetric(root *TreeNode) bool {
	return compare(root.Left, root.Right)
}
