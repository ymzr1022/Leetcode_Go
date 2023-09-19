package dp

import "fmt"

/*
给你一个 n x n 整数矩阵 grid ，请你返回 非零偏移下降路径 数字和的最小值。

非零偏移下降路径 定义为：从 grid 数组中的每一行选择一个数字，且按顺序选出来的数字中，相邻数字不在原数组的同一列。



示例 1：



输入：grid = [[1,2,3],[4,5,6],[7,8,9]]
输出：13
解释：
所有非零偏移下降路径包括：
[1,5,9], [1,5,7], [1,6,7], [1,6,8],
[2,4,8], [2,4,9], [2,6,7], [2,6,8],
[3,4,8], [3,4,9], [3,5,7], [3,5,9]
下降路径中数字和最小的是 [1,5,7] ，所以答案是 13 。
示例 2：

输入：grid = [[7]]
输出：7


提示：

n == grid.length == grid[i].length
1 <= n <= 200
-99 <= grid[i][j] <= 99
*/

func MinFallingPathSums(grid [][]int) int {
	n := len(grid)
	m := len(grid[0])
	dp := make([][]int, 2)
	for i := range dp {
		dp[i] = make([]int, m)
	}
	for i := 0; i < m; i++ {
		dp[(n-1)%2][i] = grid[n-1][i]
	}
	res := 1<<31 - 1
	for i := n - 2; i >= 0; i-- {
		for j := 0; j < m; j++ {
			minx := 1<<31 - 1
			for k := 0; k < m; k++ {
				if j != k {
					minx = min(dp[(i+1)%2][k], minx)
				}
			}
			dp[i%2][j] = grid[i][j] + minx
			if i == 0 {
				res = min(dp[i%2][j], res)
			}
		}

	}
	if n == 1 {
		for i := 0; i < m; i++ {
			res = min(dp[(n-1)%2][i], res)
		}
	}
	fmt.Printf("dp: %v\n", dp)
	return res
}
