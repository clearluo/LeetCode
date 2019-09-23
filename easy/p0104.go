package easy

import "math"

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return int(math.Max(float64(maxDepth(root.Left)), float64(maxDepth(root.Right))) + 1)
	//lDeep := maxDepth(root.Left)
	//rDeep := maxDepth(root.Right)
	//if lDeep > rDeep {
	//	return lDeep + 1
	//}
	//return rDeep + 1
}
