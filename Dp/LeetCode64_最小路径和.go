package dp

import (
	"math"
)

/**
给定一个包含非负整数的 m x n 网格 grid ，请找出一条从左上角到右下角的路径，使得路径上的数字总和为最小。

说明：每次只能向下或者向右移动一步。



示例 1：


输入：grid = [[1,3,1],[1,5,1],[4,2,1]]
输出：7
解释：因为路径 1→3→1→1→1 的总和最小。
示例 2：

输入：grid = [[1,2,3],[4,5,6]]
输出：12


提示：

m == grid.length
n == grid[i].length
1 <= m, n <= 200
0 <= grid[i][j] <= 200

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/minimum-path-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func MinPathSum(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])
	// var process func(int, int) int
	// process = func(i, j int) int {
	// 	if i == 0 && j == 0 {
	// 		return grid[0][0]
	// 	}
	// 	if i < 0 || j < 0 {
	// 		return math.MaxInt32
	// 	}
	// 	return min(process(i-1, j), process(i, j-1)) + grid[i][j]
	// }
	// return process(row-1, col-1)

	dp := make([][]int, row+1)
	for i := range dp {
		dp[i] = make([]int, col+1)
		dp[i][0] = math.MaxInt32
		for k := range dp[i] {
			dp[0][k] = math.MaxInt32
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if i == 0 && j == 0 {
				dp[i+1][j+1] = grid[i][i]
			} else {
				dp[i+1][j+1] = min(dp[i][j+1], dp[i+1][j]) + grid[i][j]
			}
		}
	}
	return dp[row][col]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
