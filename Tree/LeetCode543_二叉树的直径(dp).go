package tree

/**
给你一棵二叉树的根节点，返回该树的 直径 。

二叉树的 直径 是指树中任意两个节点之间最长路径的 长度 。这条路径可能经过也可能不经过根节点 root 。

两节点之间路径的 长度 由它们之间边数表示。



示例 1：


输入：root = [1,2,3,4,5]
输出：3
解释：3 ，取路径 [4,2,1,3] 或 [5,2,1,3] 的长度。
示例 2：

输入：root = [1,2]
输出：1


提示：

树中节点数目在范围 [1, 104] 内
-100 <= Node.val <= 100

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/diameter-of-binary-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func DiameterOfBinaryTree(root *TreeNode) int {
	ans := 0
	var process func(*TreeNode) int
	process = func(tn *TreeNode) int {
		if tn == nil {
			return -1
		}
		l := process(tn.Left)
		r := process(tn.Right)
		ans = max(ans, l+r+2)
		return max(l, r)
	}
	return process(root)
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
