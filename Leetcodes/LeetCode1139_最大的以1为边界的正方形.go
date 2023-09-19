package leetcodes

/**
给你一个由若干 0 和 1 组成的二维网格 grid，请你找出边界全部由 1 组成的最大 正方形 子网格，并返回该子网格中的元素数量。如果不存在，则返回 0。



示例 1：

输入：grid = [[1,1,1],[1,0,1],[1,1,1]]
输出：9
示例 2：

输入：grid = [[1,1,0,0]]
输出：1


提示：

1 <= grid.length <= 100
1 <= grid[0].length <= 100
grid[i][j] 为 0 或 1

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/largest-1-bordered-square
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func largest1BorderedSquare(grid [][]int) int {
	row := len(grid)
	col := len(grid[0])

	rx := make([][]int, row)
	cx := make([][]int, col)
	for i := range rx {
		rx[i] = make([]int, col+1)
	}
	for i := range cx {
		cx[i] = make([]int, row+1)
	}

	for i, row := range grid {
		for j, x := range row {
			rx[i][j+1] = rx[i][j] + x // 每行的前缀和
			cx[j][i+1] = cx[j][i] + x // 每列的前缀和
		}
	}

	// for i := 0; i < row; i++ {
	// 	for j := 0; j < col; j++ {
	// 		rx[i][j+1] = rx[i][j] + grid[i][j]
	// 		cx[j][i+1] = cx[j][i] + grid[i][j]
	// 	}
	// }

	for d := min(row, col); d >= 0; d-- {
		for i := 0; i <= row-d; i++ {
			for j := 0; j <= col-d; j++ {
				if rx[i][j+d]-rx[i][j] == d &&
					cx[j][i+d]-cx[j][i] == d &&
					rx[i+d-1][j+d]-rx[i+d-1][j] == d &&
					cx[j+d-1][i+d]-cx[j+d-1][i] == d {
					return d * d
				}
			}
		}
	}
	return 0
}
