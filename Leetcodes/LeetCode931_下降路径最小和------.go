package leetcodes

/**
给你一个 n x n 的 方形 整数数组 matrix ，请你找出并返回通过 matrix 的下降路径 的 最小和 。

下降路径 可以从第一行中的任何元素开始，并从每一行中选择一个元素。在下一行选择的元素和当前行所选元素最多相隔一列（即位于正下方或者沿对角线向左或者向右的第一个元素）。具体来说，位置 (row, col) 的下一个元素应当是 (row + 1, col - 1)、(row + 1, col) 或者 (row + 1, col + 1) 。



示例 1：



输入：matrix = [[2,1,3],[6,5,4],[7,8,9]]
输出：13
解释：如图所示，为和最小的两条下降路径
示例 2：



输入：matrix = [[-19,57],[-40,-5]]
输出：-59
解释：如图所示，为和最小的下降路径


提示：

n == matrix.length == matrix[i].length
1 <= n <= 100
-100 <= matrix[i][j] <= 100

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/minimum-falling-path-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func minFallingPathSum(matrix [][]int) (res int) {

	// sum := 0

	// var process func(int, int) int

	// process = func(row, col int) (min int) {

	// 	if col < 0 || row > len(matrix) || col >= len(matrix[0]) {
	// 		return 0
	// 	}
	// 	left := process(row+1, col-1)
	// 	down := process(row+1, col)
	// 	right := process(row+1, col+1)

	// 	return
	// }

	// x := process(0, 0)
	return
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
