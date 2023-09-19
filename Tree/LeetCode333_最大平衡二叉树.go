package tree

// type result struct {
// 	maxs int
// 	mins int
// 	size int
// }

func LargestBSTSubtree(root *TreeNode) (res int) {
	var process func(*TreeNode) (int, int, int, bool)
	process = func(tn *TreeNode) (int, int, int, bool) {
		if tn == nil {
			return 0, -1 << 31, 1 << 31, true
		}
		if tn.Left == nil && tn.Right == nil {
			return 1, tn.Val, tn.Val, true
		}
		sizeL, maxL, minL, okL := process(tn.Left)
		sizeR, maxR, minR, okR := process(tn.Right)
		if !okL || !okR {
			return 0, -1 << 31, 1 << 31, false
		}
		if root.Left != nil && maxL >= tn.Val {
			return max(sizeL, sizeR), 0, 0, false
		}
		if root.Right != nil && minR <= tn.Val {
			return max(sizeL, sizeR), 0, 0, false
		}
		if root.Left == nil {
			return 1 + sizeR, maxR, tn.Val, true
		}
		if root.Right == nil {
			return 1 + sizeL, tn.Val, minL, true
		}
		return 1 + sizeL + sizeR, minL, maxR, true
	}
	// process(root)
	x, _, _, _ := process(root)
	res = x
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
