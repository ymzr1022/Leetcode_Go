package tree

/**
小偷又发现了一个新的可行窃的地区。这个地区只有一个入口，我们称之为 root 。

除了 root 之外，每栋房子有且只有一个“父“房子与之相连。一番侦察之后，聪明的小偷意识到“这个地方的所有房屋的排列类似于一棵二叉树”。 如果 两个直接相连的房子在同一天晚上被打劫 ，房屋将自动报警。

给定二叉树的 root 。返回 在不触动警报的情况下 ，小偷能够盗取的最高金额 。



示例 1:



输入: root = [3,2,3,null,3,null,1]
输出: 7
解释: 小偷一晚能够盗取的最高金额 3 + 3 + 1 = 7
示例 2:



输入: root = [3,4,5,1,3,null,1]
输出: 9
解释: 小偷一晚能够盗取的最高金额 4 + 5 = 9


提示：

树的节点数在 [1, 104] 范围内
0 <= Node.val <= 104
通过次数282,229提交次数461,871

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/house-robber-iii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// func Rob3(root *TreeNode) int {
// 	var process func(*TreeNode) (int, int)
// 	process = func(tn *TreeNode) (se int, nse int) {
// 		if tn == nil {
// 			return 0, 0
// 		}
// 		sel, nsel := process(tn.Left)
// 		ser, nser := process(tn.Right)
// 		se = tn.Val + nsel + nser
// 		nse = max(sel, nsel) + max(ser, nser)
// 		return se, nse
// 	}
// 	se, nse := process(root)
// 	return max(se, nse)
// }

func rob(root *TreeNode) int {
	var dfs func(node *TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		left1, left2 := dfs(node.Left)
		right1, right2 := dfs(node.Right)
		return node.Val + left2 + right2, left1 + right1
	}
	x, y := dfs(root)
	return max(x, y)
}
