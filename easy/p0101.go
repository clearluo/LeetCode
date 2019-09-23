package easy

func isSymmetric(root *TreeNode) bool {
	// TODO 未完成
	if root == nil {
		return true
	}
	if root.Left != root.Right && (root.Left == nil || root.Right == nil) {
		return false
	}
	if root.Left != nil && root.Right != nil && (root.Left.Val != root.Right.Val) {
		return false
	}
	return isSymmetric(root.Left) && isSymmetric(root.Right)
}
